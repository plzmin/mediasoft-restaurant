package productservice

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"restaurant/internal/model"
	"time"
)

func (s *Service) CreateProduct(ctx context.Context,
	req *restaurant.CreateProductRequest) (*restaurant.CreateProductResponse, error) {

	product := model.Product{
		Uuid:        uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
		Weight:      req.Weight,
		Price:       req.Price,
		CreatedAt:   time.Now(),
	}

	if err := s.productRepository.Create(ctx, &product); err != nil {
		s.log.Error("failed create product %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &restaurant.CreateProductResponse{}, nil
}
