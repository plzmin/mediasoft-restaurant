package menusqlx

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"restaurant/internal/model"
	"time"
)

type MenuSqlx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *MenuSqlx {
	return &MenuSqlx{db: db}
}

func (r *MenuSqlx) Create(ctx context.Context, m *model.Menu) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	const mq = `insert into menu (uuid, on_date,opening_record_at,closing_record_at)
						values (:uuid, :on_date,:opening_record_at,:closing_record_at)`
	_, err = tx.NamedExecContext(ctx, mq, m)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errors.Join(err, errRollback)
		}
	}

	mpq := `insert into menu_product(menu_uuid, product_uuid) values ($1, $2)`
	for _, pl := range [][]string{m.Salads, m.Garnishes, m.Meats, m.Soups, m.Drinks, m.Desserts} {
		for _, p := range pl {
			_, err := tx.ExecContext(ctx, mpq, m.Uuid, p)
			if err != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					return errors.Join(err, errRollback)
				}
			}
		}
	}

	return tx.Commit()
}

func (r *MenuSqlx) Get(ctx context.Context, time time.Time) (*model.Menu, []*model.Product, error) {
	const mq = "select * from menu  where date_trunc('day',on_date) =  date_trunc('day',$1::timestamp)"
	menu := new(model.Menu)
	if err := r.db.GetContext(ctx, menu, mq, time); err != nil {
		return nil, nil, err
	}

	const productq = `select p.uuid, p.name, p.description, p.type, p.weight, p.price, p.created_at
						from products p join menu_product mp on mp.product_uuid=p.uuid where mp.menu_uuid=$1`

	var products []*model.Product
	if err := r.db.SelectContext(ctx, &products, productq, menu.Uuid); err != nil {
		return nil, nil, err
	}

	return menu, products, nil
}
