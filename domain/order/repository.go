package order

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/cart"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"gorm.io/gorm"
)

type Repository interface {
	Create(order *Order) error
	GetByOrderID(id string) (resultOrder []Order, err error)
	GetAll() (orderResult []Order, err error)
	GetLatestOrderID() (resultOrder *Order, err error)
	Upsert(order *Order) error

	// cart domain repository
	GetAllCart() (resultCart []cart.Cart, err error)

	//product domain repository
	GetProductByID(id uint64) (resultProduct *products.Products, err error)
	UpdateProduct(product *products.Products) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}

// GetLatestOrderID implements Repository
func (repo *repository) GetLatestOrderID() (resultOrder *Order, err error) {
	res := repo.db.Order("order_id desc").First(&resultOrder).Limit(1)
	return resultOrder, res.Error
}

// Create implements Repository
func (repo *repository) Create(order *Order) error {
	return repo.db.Create(order).Error
}

// GetAll implements Repository
func (repo *repository) GetAll() (orderResult []Order, err error) {
	res := repo.db.Find(&orderResult).Order("order_id desc")
	return orderResult, res.Error
}

// GetByID implements Repository
func (repo *repository) GetByOrderID(id string) (resultOrder []Order, err error) {
	res := repo.db.Where("order_id = ?", id).Find(&resultOrder)
	return resultOrder, res.Error
}

// Update implements Repository
func (repo *repository) Upsert(order *Order) error {
	return repo.db.Save(order).Error
}

//product domain repository
// GetProductByID implements Repository
func (repo *repository) GetProductByID(id uint64) (resultProduct *products.Products, err error) {
	res := repo.db.Where("ID = ?", id).Find(&resultProduct)
	return resultProduct, res.Error
}

// UpdateProduct implements Repository
func (repo *repository) UpdateProduct(product *products.Products) error {
	return repo.db.Save(product).Error
}

// cart domain repository
// GetAllCart implements Repository
func (repo *repository) GetAllCart() (resultCart []cart.Cart, err error) {
	res := repo.db.Find(&resultCart)
	return resultCart, res.Error
}
