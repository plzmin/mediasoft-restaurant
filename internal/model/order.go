package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Uuid      uuid.UUID    `json:"uuid" db:"uuid"`
	UserUuid  uuid.UUID    `json:"userUuid" db:"user_uuid"`
	Salads    []*OrderItem `json:"salads,omitempty"`
	Garnishes []*OrderItem `json:"garnishes,omitempty"`
	Meats     []*OrderItem `json:"meats,omitempty"`
	Soups     []*OrderItem `json:"soups,omitempty"`
	Drinks    []*OrderItem `json:"drinks,omitempty"`
	Desserts  []*OrderItem `json:"desserts,omitempty"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
}

type OrderItem struct {
	Count       int32     `json:"count" db:"count"`
	ProductUuid uuid.UUID `json:"product_uuid" db:"product_uuid"`
}
