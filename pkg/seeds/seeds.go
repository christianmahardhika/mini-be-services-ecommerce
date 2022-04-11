package seeds

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/pkg/seed"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
		{
			Name: "Insert Products",
			Run: func(db *gorm.DB) error {
				var err error
				CreateProducts(db, "Shampo Sunslik", 100000, "GET20OFF", "PERCENTAGE", 20, 10)
				CreateProducts(db, "Sabun Lifeboy", 100000, "GET30OFF", "PERCENTAGE", 30, 10)
				return err
			},
		},
	}
}
