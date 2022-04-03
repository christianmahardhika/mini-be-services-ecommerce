package seeds

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/pkg/seed"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "Insert Products",
			Run: func(db *gorm.DB) error {
				var err error
				CreateProducts(db, "Product 1", 100000, "BELI2GRATIS1", 10)
				CreateProducts(db, "Product 2", 200000, "BELI2GRATIS1", 10)
				CreateProducts(db, "Product 3", 300000, "BELI3GRATIS1", 10)
				CreateProducts(db, "Product 4", 400000, "DISKON10%", 10)
				CreateProducts(db, "Product 5", 500000, "DISKON10", 10)

				return err
			},
		},
	}
}
