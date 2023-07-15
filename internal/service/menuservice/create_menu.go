package menuservice

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"restaurant/internal/model"
	"time"
)

func (s *Service) CreateMenu(ctx context.Context,
	req *restaurant.CreateMenuRequest) (*restaurant.CreateMenuResponse, error) {
	if err := req.ValidateAll(); err != nil {
		s.log.Warn("not valid CreateMenuRequest %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	salads, err := strToProduct(req.Salads)
	if err != nil {
		s.log.Error("failed to parse saladsUUID %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	garnishes, err := strToProduct(req.Garnishes)
	if err != nil {
		s.log.Error("failed to parse garnishesUUID %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	meats, err := strToProduct(req.Meats)
	if err != nil {
		s.log.Error("failed to parse meatsUUID %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	soups, err := strToProduct(req.Soups)
	if err != nil {
		s.log.Error("failed to parse soupsUUID %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	drinks, err := strToProduct(req.Drinks)
	if err != nil {
		s.log.Error("failed to parse drinksUUID %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	desserts, err := strToProduct(req.Desserts)
	if err != nil {
		s.log.Error("failed to parse dessertsUUID %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	menu := model.Menu{
		Uuid:            uuid.New(),
		OnDate:          req.OnDate.AsTime(),
		OpeningRecordAt: req.OpeningRecordAt.AsTime(),
		ClosingRecordAt: req.ClosingRecordAt.AsTime(),
		Salads:          salads,
		Garnishes:       garnishes,
		Meats:           meats,
		Soups:           soups,
		Drinks:          drinks,
		Desserts:        desserts,
		CreatedAt:       time.Now(),
	}

	if err = s.menuRepository.Create(ctx, &menu); err != nil {
		s.log.Error("failed create menu %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &restaurant.CreateMenuResponse{}, nil
}

func strToProduct(str []string) ([]*model.Product, error) {
	var data []*model.Product
	for _, p := range str {
		id, err := uuid.Parse(p)
		if err != nil {
			return nil, err
		}
		data = append(data, &model.Product{Uuid: id})
	}

	return data, nil
}
