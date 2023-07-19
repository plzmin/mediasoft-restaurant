package model

import (
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"time"
)

type Product struct {
	Uuid        uuid.UUID              `db:"uuid"`
	Name        string                 `db:"name"`
	Description string                 `db:"description"`
	Type        restaurant.ProductType `db:"type"`
	Weight      int32                  `db:"weight"`
	Price       float64                `db:"price"`
	CreatedAt   time.Time              `db:"created_at"`
}
