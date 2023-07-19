package menuservice

import (
	"context"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"restaurant/internal/model"
)

func (s *Service) GetMenu(ctx context.Context,
	req *restaurant.GetMenuRequest) (*restaurant.GetMenuResponse, error) {
	menu, products, err := s.menuRepository.Get(ctx, req.OnDate.AsTime())
	if err != nil {
		s.log.Error("failed get menu %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	m := restaurant.Menu{
		Uuid:            menu.Uuid.String(),
		OnDate:          timestamppb.New(menu.OnDate),
		OpeningRecordAt: timestamppb.New(menu.OpeningRecordAt),
		ClosingRecordAt: timestamppb.New(menu.ClosingRecordAt),
		CreatedAt:       timestamppb.New(menu.CreatedAt),
	}

	for _, product := range products {
		switch product.Type {
		case restaurant.ProductType_PRODUCT_TYPE_SALAD:
			m.Salads = append(m.Salads, appendProduct(product))
		case restaurant.ProductType_PRODUCT_TYPE_GARNISH:
			m.Garnishes = append(m.Garnishes, appendProduct(product))
		case restaurant.ProductType_PRODUCT_TYPE_MEAT:
			m.Meats = append(m.Meats, appendProduct(product))
		case restaurant.ProductType_PRODUCT_TYPE_SOUP:
			m.Soups = append(m.Soups, appendProduct(product))
		case restaurant.ProductType_PRODUCT_TYPE_DRINK:
			m.Drinks = append(m.Drinks, appendProduct(product))
		case restaurant.ProductType_PRODUCT_TYPE_DESSERT:
			m.Desserts = append(m.Desserts, appendProduct(product))
		}
	}
	return &restaurant.GetMenuResponse{Menu: &m}, nil
}

func appendProduct(product *model.Product) *restaurant.Product {
	p := &restaurant.Product{
		Uuid:        product.Uuid.String(),
		Name:        product.Name,
		Description: product.Description,
		Type:        product.Type,
		Weight:      product.Weight,
		Price:       product.Price,
		CreatedAt:   timestamppb.New(product.CreatedAt),
	}
	return p
}
