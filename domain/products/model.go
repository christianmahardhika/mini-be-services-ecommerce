package products

import "time"

type Products struct {
	ID        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Promo     string    `json:"promo"`
	Stock     int64     `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
