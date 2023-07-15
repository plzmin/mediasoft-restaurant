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
	menu, err := s.menuRepository.Get(ctx, req.OnDate.AsTime())
	if err != nil {
		s.log.Error("failed get menu %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	m := restaurant.Menu{
		Uuid:            menu.Uuid.String(),
		OnDate:          timestamppb.New(menu.OnDate),
		OpeningRecordAt: timestamppb.New(menu.OpeningRecordAt),
		ClosingRecordAt: timestamppb.New(menu.ClosingRecordAt),
		Salads:          modelToRestaurant(menu.Salads),
		Garnishes:       modelToRestaurant(menu.Garnishes),
		Meats:           modelToRestaurant(menu.Meats),
		Soups:           modelToRestaurant(menu.Soups),
		Drinks:          modelToRestaurant(menu.Drinks),
		Desserts:        modelToRestaurant(menu.Desserts),
		CreatedAt:       timestamppb.New(menu.CreatedAt),
	}
	return &restaurant.GetMenuResponse{Menu: &m}, nil
}

func modelToRestaurant(pl []*model.Product) []*restaurant.Product {
	var rpl []*restaurant.Product
	for _, p := range pl {
		rp := restaurant.Product{
			Uuid:        p.Uuid.String(),
			Name:        p.Name,
			Description: p.Description,
			Type:        restaurant.ProductType(p.Type),
			Weight:      p.Weight,
			Price:       p.Price,
			CreatedAt:   timestamppb.New(p.CreatedAt),
		}
		rpl = append(rpl, &rp)
	}
	return rpl
}
