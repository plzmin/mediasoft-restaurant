package menurepository

import (
	"context"
	"restaurant/internal/model"
	"time"
)

//go:generate mockery --all

type MenuRepository interface {
	Create(ctx context.Context, menu *model.Menu) error
	Get(ctx context.Context, time time.Time) (*model.Menu, []*model.Product, error)
}
