package order

import (
	"time"
)

type OrderDetail struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	OrderID   string    `gorm:"not null" json:"orderId"`
	ProductID string    `gorm:"not null" json:"productId"`
	Quantity  int64     `gorm:"not null" json:"quantity"`
	Price     int64     `gorm:"not null" json:"price"`
	PromoCode string    `gorm:"not null" json:"promoCode"`
	Total     int64     `gorm:"not null" json:"total"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Order struct {
	ID        string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    string    `json:"userId"`
	TotalCart int64     `json:"totalCart" gorm:"null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type InvoiceOrder struct {
	Order       Order
	OrderDetail []OrderDetail
}
