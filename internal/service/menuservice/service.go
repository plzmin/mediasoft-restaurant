package menuservice

import (
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"restaurant/internal/repository/menurepository"
	"restaurant/pkg/logger"
)

type Service struct {
	log            *logger.Logger
	menuRepository menurepository.MenuRepository
	restaurant.UnimplementedMenuServiceServer
}

func New(log *logger.Logger, menuRepository menurepository.MenuRepository) *Service {
	return &Service{
		log:            log,
		menuRepository: menuRepository,
	}
}
