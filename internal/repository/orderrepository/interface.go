package orderrepository

import (
	"context"
	"restaurant/internal/model"
	"time"
)

type OrderRepository interface {
	Create(model *model.Order) error
	Get(ctx context.Context, time time.Time) ([]*model.TotalOrder, error)
}
