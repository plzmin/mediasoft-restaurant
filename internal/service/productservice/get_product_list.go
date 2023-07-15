package productservice

import (
	"context"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetProductList(ctx context.Context, req *restaurant.GetProductListRequest) (*restaurant.GetProductListResponse, error) {
	list, err := s.productRepository.List(ctx)
	if err != nil {
		s.log.Error("failed get productList %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	var data []*restaurant.Product

	for _, p := range list {
		data = append(data, &restaurant.Product{
			Uuid:        p.Uuid.String(),
			Name:        p.Name,
			Description: p.Description,
			Type:        restaurant.ProductType(p.Type),
			Weight:      p.Weight,
			Price:       p.Price,
			CreatedAt:   timestamppb.New(p.CreatedAt),
		})
	}

	return &restaurant.GetProductListResponse{Result: data}, nil
}
