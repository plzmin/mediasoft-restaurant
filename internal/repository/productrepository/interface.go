package productrepository

import (
	"context"
	"restaurant/internal/model"
)

//go:generate mockery --all

type ProductRepository interface {
	Create(ctx context.Context, p *model.Product) error
	List(ctx context.Context) ([]*model.Product, error)
}
