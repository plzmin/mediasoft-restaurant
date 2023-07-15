package client

import (
	"context"
	"fmt"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"restaurant/internal/config"
)

type CustomerClient struct {
	office customer.OfficeServiceClient
	user   customer.UserServiceClient
}

func New(cfg config.Config) (*CustomerClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.CustomerGRPC.IP, cfg.CustomerGRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	office := customer.NewOfficeServiceClient(conn)
	user := customer.NewUserServiceClient(conn)

	return &CustomerClient{
		office: office,
		user:   user,
	}, err
}

func (c *CustomerClient) Close() error {
	return c.Close()
}

func (c *CustomerClient) GetUsers(ctx context.Context) (*customer.GetUserListResponse, error) {
	res, err := c.user.GetUserList(ctx, &customer.GetUserListRequest{})
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *CustomerClient) GetOffices(ctx context.Context) (*customer.GetOfficeListResponse, error) {
	res, err := c.office.GetOfficeList(ctx, &customer.GetOfficeListRequest{})
	if err != nil {
		return nil, err
	}
	return res, err
}
