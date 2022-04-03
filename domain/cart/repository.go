package cart

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Cart, error)
	UpdateInsert(cart *Cart) error
	Delete(cartID int64) error
	DeleteAll() error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}

// DeleteAll implements Repository
func (repo *repository) DeleteAll() error {
	res := repo.db.Where("id > ?", 0).Delete(&Cart{})
	return res.Error
}

// UpdateInsert implements Repository
func (repo *repository) UpdateInsert(cart *Cart) error {
	return repo.db.Save(cart).Error
}

// Delete implements Repository
func (repo *repository) Delete(cartID int64) error {
	res := repo.db.Delete(&Cart{}, cartID)
	return res.Error
}

// GetAll implements Repository
func (repo *repository) GetAll() (cartResult []Cart, err error) {
	res := repo.db.Find(&cartResult)
	return cartResult, res.Error
}
