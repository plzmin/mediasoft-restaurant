package orderservice

import (
	"context"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Service) GetUpToDateOrderList(ctx context.Context,
	req *restaurant.GetUpToDateOrderListRequest) (*restaurant.GetUpToDateOrderListResponse, error) {
	orders, err := s.orderRepository.Get(ctx, time.Now().Add(-24*time.Hour))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	users, err := s.client.GetUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	offices, err := s.client.GetOffices(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var totalOrders []*restaurant.Order
	for _, order := range orders {
		totalOrders = append(totalOrders, &restaurant.Order{
			ProductId:   order.ProductUuid.String(),
			ProductName: order.ProductName,
			Count:       order.Count,
		})
	}

	var totalOrdersByCompany []*restaurant.OrdersByOffice
	for _, office := range offices.Result {
		var orderByOffice []*restaurant.Order
		for _, user := range users.Result {
			if user.OfficeUuid == office.Uuid {
				for _, order := range orders {
					if order.UserUuid.String() == user.Uuid {
						orderByOffice = append(orderByOffice, &restaurant.Order{
							ProductId:   order.ProductUuid.String(),
							ProductName: order.ProductName,
							Count:       order.Count,
						})
					}
				}
			}

		}
		totalOrdersByCompany = append(totalOrdersByCompany, &restaurant.OrdersByOffice{
			OfficeUuid:    office.Uuid,
			OfficeName:    office.Name,
			OfficeAddress: office.Address,
			Result:        orderByOffice,
		})
	}

	return &restaurant.GetUpToDateOrderListResponse{
		TotalOrders:          totalOrders,
		TotalOrdersByCompany: totalOrdersByCompany,
	}, nil
}
