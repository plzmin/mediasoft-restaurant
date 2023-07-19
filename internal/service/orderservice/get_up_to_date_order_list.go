package orderservice

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"restaurant/internal/model"
	"time"
)

func (s *Service) GetUpToDateOrderList(ctx context.Context,
	req *restaurant.GetUpToDateOrderListRequest) (*restaurant.GetUpToDateOrderListResponse, error) {
	orders, err := s.orderRepository.Get(ctx, time.Now()) //.Add(-24*time.Hour))
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

	products := make(map[uuid.UUID]int64)
	for _, order := range orders {
		products[order.ProductUuid] += order.Count
	}

	var totalOrders []*restaurant.Order
	for id, count := range products {
		for _, o := range orders {
			if id == o.ProductUuid {
				totalOrders = append(totalOrders, &restaurant.Order{
					ProductId:   o.ProductUuid.String(),
					ProductName: o.ProductName,
					Count:       count,
				})
				break
			}
		}
	}

	var orderByOffices []*model.OrderByOffice
	for _, office := range offices.Result {
		orderByOffice := make(map[uuid.UUID]int64)
		for _, user := range users.Result {
			if user.OfficeUuid == office.Uuid {
				for _, o := range orders {
					if o.UserUuid.String() == user.Uuid {
						orderByOffice[o.ProductUuid] += o.Count
					}
				}
			}
		}
		orderByOffices = append(orderByOffices, &model.OrderByOffice{
			OfficeUuid:    office.Uuid,
			OfficeName:    office.Name,
			OfficeAddress: office.Address,
			Result:        orderByOffice,
		})
	}

	var totalOrdersByCompany []*restaurant.OrdersByOffice
	for _, order := range orderByOffices {
		var orderByOffice []*restaurant.Order
		for id, count := range order.Result {
			for _, o := range orders {
				if o.ProductUuid == id {
					orderByOffice = append(orderByOffice, &restaurant.Order{
						ProductId:   o.ProductUuid.String(),
						ProductName: o.ProductName,
						Count:       count,
					})
					break
				}
			}
		}
		if orderByOffice != nil {
			totalOrdersByCompany = append(totalOrdersByCompany, &restaurant.OrdersByOffice{
				OfficeUuid:    order.OfficeUuid,
				OfficeName:    order.OfficeName,
				OfficeAddress: order.OfficeAddress,
				Result:        orderByOffice,
			})
		}
	}

	return &restaurant.GetUpToDateOrderListResponse{
		TotalOrders:          totalOrders,
		TotalOrdersByCompany: totalOrdersByCompany,
	}, nil
}
