package productrepository

import (
	"context"
	"restaurant/internal/model"
)

type ProductRepository interface {
	Create(ctx context.Context, p *model.Product) error
	List(ctx context.Context) ([]*model.Product, error)
}
