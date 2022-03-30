package products

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]*Products, error)
	GetByID(id uint64) (*Products, error)
	Create(p *Products) error
	Update(p *Products) error
	Delete(id uint64) error
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Create implements Repository
func (repo *repository) Create(p *Products) error {
	panic("unimplemented")
}

// Delete implements Repository
func (repo *repository) Delete(id uint64) error {
	panic("unimplemented")
}

// GetAll implements Repository
func (repo *repository) GetAll() ([]*Products, error) {
	panic("unimplemented")
}

// GetByID implements Repository
func (repo *repository) GetByID(id uint64) (*Products, error) {
	panic("unimplemented")
}

// Update implements Repository
func (repo *repository) Update(p *Products) error {
	panic("unimplemented")
}
