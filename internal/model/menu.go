package model

import (
	"github.com/google/uuid"
	"time"
)

type Menu struct {
	Uuid            uuid.UUID `json:"uuid" db:"uuid"`
	OnDate          time.Time `json:"on_date" db:"on_date"`
	OpeningRecordAt time.Time `json:"opening_record_at" db:"opening_record_at"`
	ClosingRecordAt time.Time `json:"closing_record_at" db:"closing_record_at"`
	Salads          []string  `json:"salads"`
	Garnishes       []string  `json:"garnishes"`
	Meats           []string  `json:"meats"`
	Soups           []string  `json:"soups"`
	Drinks          []string  `json:"drinks"`
	Desserts        []string  `json:"desserts"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type MenuProduct struct {
	MenuUuid    uuid.UUID `json:"menu_uuid" db:"menu_uuid"`
	ProductUuid uuid.UUID `json:"product_uuid" db:"product_uuid"`
}
