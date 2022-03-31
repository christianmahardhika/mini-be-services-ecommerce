package seeds

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"gorm.io/gorm"
)

func CreateProducts(db *gorm.DB, name string, price int64, promo string, qty int64) error {
	return db.Create(&products.Products{Name: name, Price: price, Promo: promo, Quantity: qty}).Error
}
