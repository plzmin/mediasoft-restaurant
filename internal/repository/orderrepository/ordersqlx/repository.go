package ordersqlx

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"restaurant/internal/model"
	"time"
)

type OrderSqlx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *OrderSqlx {
	return &OrderSqlx{db: db}
}

func (r *OrderSqlx) Create(order *model.Order) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	const q = `insert into orders (uuid, user_uuid) values(:uuid, :user_uuid)`
	_, err = tx.NamedExec(q, order)
	if err != nil {
		errRollback := tx.Rollback()
		if err != nil {
			return errors.Join(err, errRollback)
		}
	}

	orderItemq := `insert into order_item(order_uuid, count, product_uuid) values ($1, $2, $3)`
	for _, orderItems := range [][]*model.OrderItem{order.Salads, order.Soups, order.Drinks, order.Desserts, order.Meats, order.Garnishes} {
		for _, orderItem := range orderItems {
			_, err := tx.Exec(orderItemq, order.Uuid, orderItem.Count, orderItem.ProductUuid)
			if err != nil {
				errRollback := tx.Rollback()
				if err != nil {
					return errors.Join(err, errRollback)
				}
			}
		}
	}
	return tx.Commit()
}

func (r *OrderSqlx) Get(ctx context.Context, time time.Time) ([]*model.TotalOrder, error) {
	const q = `select orders.*, order_item.count, order_item.product_uuid, products.name from orders 
    	join order_item on orders.uuid = order_item.order_uuid
    	join products on order_item.product_uuid = products.uuid where date_trunc('day',orders.created_at) = date_trunc('day', $1::timestamp)`
	var list []*model.TotalOrder
	err := r.db.SelectContext(ctx, &list, q, time)
	return list, err
}
