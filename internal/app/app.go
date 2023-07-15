package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"os/signal"
	"restaurant/internal/bootstrap"
	"restaurant/internal/client"
	"restaurant/internal/config"
	"restaurant/internal/kafka"
	"restaurant/internal/repository/menurepository/menusqlx"
	"restaurant/internal/repository/orderrepository/ordersqlx"
	"restaurant/internal/repository/productrepository/productsqlx"
	"restaurant/internal/service/menuservice"
	"restaurant/internal/service/orderservice"
	"restaurant/internal/service/productservice"
	"restaurant/pkg/logger"
	"syscall"
)

func Run(log *logger.Logger, cfg config.Config) error {

	s := grpc.NewServer()
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	go runGRPCServer(log, cfg, s)
	go runHTTPServer(log, ctx, cfg, mux)

	gracefulShutDown(log, s, cancel)

	return nil
}

func runGRPCServer(log *logger.Logger, cfg config.Config, s *grpc.Server) {
	db, err := bootstrap.InitSqlxDB(cfg)
	if err != nil {
		log.Fatal("failed to init db conn %v", err.Error())
	}

	consumer, err := kafka.New(cfg.Kafka, log, ordersqlx.New(db))
	if err != nil {
		log.Fatal("failed get conn kafka %v", err.Error())
	}
	defer func(consumer *kafka.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Fatal("failed get conn kafka %v", err.Error())
		}
	}(consumer)

	topic := "order"
	consumer.Consume(topic)

	с, err := client.New(cfg)
	if err != nil {
		log.Fatal("failed get conn customer %v", err.Error())
	}
	defer func(с *client.CustomerClient) {
		err := с.Close()
		if err != nil {
			log.Fatal("failed get conn customer %v", err.Error())
		}
	}(с)

	productService := productservice.New(log, productsqlx.New(db))
	menuService := menuservice.New(log, menusqlx.New(db))
	orderService := orderservice.New(log, с, ordersqlx.New(db))
	restaurant.RegisterProductServiceServer(s, productService)
	restaurant.RegisterMenuServiceServer(s, menuService)
	restaurant.RegisterOrderServiceServer(s, orderService)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.IP, cfg.GRPC.Port))
	if err != nil {
		log.Fatal("failed to listen tcp %s:%d, %v", cfg.GRPC.IP, cfg.GRPC.Port, err)
	}

	log.Info("starting listening grpc server at %s:%d", cfg.GRPC.IP, cfg.GRPC.Port)
	if err = s.Serve(l); err != nil {
		log.Fatal("error service grpc server %v", err)
	}

}

func runHTTPServer(log *logger.Logger, ctx context.Context, cfg config.Config, mux *runtime.ServeMux) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endPoint := fmt.Sprintf("%s:%d", cfg.HTTP.IP, cfg.GRPC.Port)

	if err := restaurant.RegisterMenuServiceHandlerFromEndpoint(ctx, mux, endPoint, dialOptions); err != nil {
		log.Fatal("failed to register menu service %v", err)
	}

	if err := restaurant.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, endPoint, dialOptions); err != nil {
		log.Fatal("failed to register order service %v", err)
	}

	if err := restaurant.RegisterProductServiceHandlerFromEndpoint(ctx, mux, endPoint, dialOptions); err != nil {
		log.Fatal("failed to register order service %v", err)
	}

	log.Info("starting listening http server at %s:%d", cfg.HTTP.IP, cfg.HTTP.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.HTTP.IP, cfg.HTTP.Port), mux); err != nil {
		log.Fatal("error service http server %v", err)
	}

}

func gracefulShutDown(log *logger.Logger, s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)
	sig := <-ch
	log.Info("Received shutdown signal: %v -  Graceful shutdown done", sig)
	s.GracefulStop()
	cancel()
}
