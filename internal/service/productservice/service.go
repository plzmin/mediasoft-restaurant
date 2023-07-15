package productservice

import (
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"restaurant/internal/repository/productrepository"
	"restaurant/pkg/logger"
)

type Service struct {
	log               *logger.Logger
	productRepository productrepository.ProductRepository
	restaurant.UnimplementedProductServiceServer
}

func New(log *logger.Logger, productRepository productrepository.ProductRepository) *Service {
	return &Service{
		log:               log,
		productRepository: productRepository,
	}
}
