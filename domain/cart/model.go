package cart

import (
	"time"
)

type Cart struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID    string    `gorm:"null" json:"userId"`
	ProductID string    `gorm:"not null" json:"productId"`
	Quantity  int64     `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
