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
				CreateProducts(db, "Product 1", 100000, "Beli 2 Gratis 1", 10)
				CreateProducts(db, "Product 2", 200000, "Beli 3 Gratis 1", 10)
				CreateProducts(db, "Product 3", 300000, "Beli 4 Gratis 1", 10)
				CreateProducts(db, "Product 4", 400000, "Beli 5 Gratis 1", 10)
				CreateProducts(db, "Product 5", 500000, "Beli 2 Gratis 1", 10)

				return err
			},
		},
	}
}
