package orderservice

import (
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"restaurant/internal/client"
	"restaurant/internal/repository/orderrepository"
	"restaurant/pkg/logger"
)

type Service struct {
	log             *logger.Logger
	client          *client.CustomerClient
	orderRepository orderrepository.OrderRepository
	restaurant.UnimplementedOrderServiceServer
}

func New(log *logger.Logger,
	client *client.CustomerClient,
	orderRepository orderrepository.OrderRepository) *Service {
	return &Service{
		log:             log,
		client:          client,
		orderRepository: orderRepository,
	}
}
