package menusqlx

import (
	"context"
	"fmt"
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
	const mq = `insert into menu (uuid, on_date,opening_record_at,closing_record_at,created_at)
						values (:uuid, :on_date,:opening_record_at,:closing_record_at, :created_at)`
	_, err = tx.NamedExec(mq, m)
	if err != nil {
		tx.Rollback()
		return err
	}

	menuProductq := `insert into menu_product(menu_uuid, product_uuid) values `
	for _, pl := range [][]*model.Product{m.Salads, m.Garnishes, m.Meats, m.Soups, m.Drinks, m.Desserts} {
		for _, p := range pl {
			menuProductq += fmt.Sprintf("($%s,$%s),", m.Uuid, p.Uuid)
		}
	}
	_, err = tx.Exec(menuProductq[:len(menuProductq)-1])
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return err
}

func (r *MenuSqlx) Get(ctx context.Context, time time.Time) (*model.Menu, error) {
	const mq = "select * from menu  where date_trunc('day',on_date) =  date_trunc('day',$1::timestamp)"
	m := new(model.Menu)
	if err := r.db.GetContext(ctx, m, mq, time); err != nil {
		return nil, err
	}

	const mpq = "select * from menu_product where menu_uuid = $1"
	var mpl []*model.MenuProduct
	if err := r.db.SelectContext(ctx, &mpl, mpq, m.Uuid); err != nil {
		return nil, err
	}

	const prq = "select * from products where uuid = $1"
	for _, mp := range mpl {
		pr := new(model.Product)
		if err := r.db.GetContext(ctx, pr, prq, mp.ProductUuid); err != nil {
			return nil, err
		}
		switch pr.Type {
		case model.PRODUCT_TYPE_SALAD:
			m.Salads = append(m.Salads, pr)
		case model.PRODUCT_TYPE_GARNISH:
			m.Garnishes = append(m.Garnishes, pr)
		case model.PRODUCT_TYPE_MEAT:
			m.Meats = append(m.Meats, pr)
		case model.PRODUCT_TYPE_SOUP:
			m.Soups = append(m.Soups, pr)
		case model.PRODUCT_TYPE_DRINK:
			m.Drinks = append(m.Drinks, pr)
		case model.PRODUCT_TYPE_DESSERT:
			m.Desserts = append(m.Desserts, pr)
		}
	}

	return m, nil
}
