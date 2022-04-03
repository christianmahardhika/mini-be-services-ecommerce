package products

import "gorm.io/gorm"

type Repository interface {
	FindByKeyword(keyword string) ([]*Products, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// GetByKeyword implements Repository
func (repo *repository) FindByKeyword(keyword string) (products []*Products, err error) {

	results := repo.db.Where("name LIKE ?", "%"+keyword+"%").Find(&products)
	return products, results.Error
}
