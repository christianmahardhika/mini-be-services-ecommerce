package products

import (
	"time"
)

type Products struct {
	ID        string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `json:"name" gorm:"uniqueIndex"`
	Price     int64     `json:"price"`
	PromoID   string    `json:"promoId"`
	Stock     int64     `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Promo struct {
	ID        string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	PromoCode string    `json:"promoCode" gorm:"uniqueIndex"`
	PromoType string    `json:"promoType"`
	Discount  int64     `json:"discount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
