package seeds

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"gorm.io/gorm"
)

func CreateProducts(db *gorm.DB, name string, price int64, promoCode string, promoType string, discount int64, qty int64) error {

	promo := products.Promo{
		PromoCode: promoCode,
		PromoType: promoType,
		Discount:  discount,
	}
	res := db.Save(&promo)
	if res.Error != nil {
		return res.Error
	}
	product := products.Products{
		Name:    name,
		Price:   price,
		PromoID: promo.ID,
		Stock:   qty,
	}
	return db.Save(&product).Error
}
