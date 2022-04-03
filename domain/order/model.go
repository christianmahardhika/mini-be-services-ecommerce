package order

import "time"

type Order struct {
	ID        uint64    `json:"id"`
	OrderID   string    `json:"orderId"`
	ProductID uint64    `json:"productId"`
	Quantity  int64     `json:"quantity"`
	Price     int64     `json:"price"`
	PromoCode string    `json:"promoCode"`
	Total     int64     `json:"total"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
