package menuservice

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"restaurant/internal/model"
)

func (s *Service) CreateMenu(ctx context.Context,
	req *restaurant.CreateMenuRequest) (*restaurant.CreateMenuResponse, error) {
	if err := req.ValidateAll(); err != nil {
		s.log.Warn("not valid CreateMenuRequest %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	menu := model.Menu{
		Uuid:            uuid.New(),
		OnDate:          req.OnDate.AsTime(),
		OpeningRecordAt: req.OpeningRecordAt.AsTime(),
		ClosingRecordAt: req.ClosingRecordAt.AsTime(),
		Salads:          req.Salads,
		Garnishes:       req.Garnishes,
		Meats:           req.Meats,
		Soups:           req.Soups,
		Drinks:          req.Drinks,
		Desserts:        req.Desserts,
	}

	if err := s.menuRepository.Create(ctx, &menu); err != nil {
		s.log.Error("failed create menu %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &restaurant.CreateMenuResponse{}, nil
}
