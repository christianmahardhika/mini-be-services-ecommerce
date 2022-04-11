package order

import (
	"time"

	"github.com/google/uuid"
)

type OrderDetail struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key",json:"id"`
	OrderID   uuid.UUID `gorm:"not null",json:"orderId"`
	ProductID uuid.UUID `gorm:"not null",json:"productId"`
	Quantity  int64     `gorm:"not null",json:"quantity"`
	Price     int64     `gorm:"not null",json:"price"`
	PromoCode string    `gorm:"not null",json:"promoCode"`
	Total     int64     `gorm:"not null",json:"total"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Order struct {
	ID        uuid.UUID `json:"id",gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID `json:"userId"`
	TotalCart int64     `json:"totalCart",gorm:"null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
