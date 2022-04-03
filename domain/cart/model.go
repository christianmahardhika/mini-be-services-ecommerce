package cart

import "time"

type Cart struct {
	ID        int64     `gorm:"primary_key"`
	ProductID int64     `gorm:"not null"`
	Quantity  int64     `gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
