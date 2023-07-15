package model

import (
	"github.com/google/uuid"
	"time"
)

type ProductType int32

const (
	PRODUCT_TYPE_UNSPECIFIED ProductType = iota
	PRODUCT_TYPE_SALAD
	PRODUCT_TYPE_GARNISH
	PRODUCT_TYPE_MEAT
	PRODUCT_TYPE_SOUP
	PRODUCT_TYPE_DRINK
	PRODUCT_TYPE_DESSERT
)

type Product struct {
	Uuid        uuid.UUID   `json:"uuid" db:"uuid"`
	Name        string      `json:"name" db:"name"`
	Description string      `json:"description" db:"description"`
	Type        ProductType `json:"type" db:"type"`
	Weight      int32       `json:"weight" db:"weight"`
	Price       float64     `json:"price" db:"price"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
}
