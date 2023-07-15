package model

import (
	"github.com/google/uuid"
	"time"
)

type TotalOrder struct {
	Uuid        uuid.UUID `db:"uuid"`
	UserUuid    uuid.UUID `db:"user_uuid"`
	CreatedAt   time.Time `db:"created_at"`
	Count       int64     `db:"count"`
	ProductUuid uuid.UUID `db:"product_uuid"`
	ProductName string    `db:"name"`
}
