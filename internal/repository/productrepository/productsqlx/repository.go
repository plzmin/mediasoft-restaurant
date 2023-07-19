package productsqlx

import (
	"context"
	"github.com/jmoiron/sqlx"
	"restaurant/internal/model"
)

type ProductSqlx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *ProductSqlx {
	return &ProductSqlx{db: db}
}

func (r *ProductSqlx) Create(ctx context.Context, product *model.Product) error {
	const q = `insert into products (uuid, name, description, type, weight, price) 
					values (:uuid, :name,:description,:type, :weight, :price)`
	_, err := r.db.NamedExecContext(ctx, q, product)
	return err
}

func (r *ProductSqlx) List(ctx context.Context) ([]*model.Product, error) {
	const q = `select * from products`
	var list []*model.Product
	err := r.db.SelectContext(ctx, &list, q)
	return list, err
}
