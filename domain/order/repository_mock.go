package order

import (
	"errors"

	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/cart"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	Mock mock.Mock
}

func (repository *OrderRepositoryMock) CreateOrderDetail(orderDetail *OrderDetail) error {
	arguments := repository.Mock.Called(orderDetail)
	if arguments.Get(0) == nil {
		return errors.New("error CreateOrderDetail")
	} else {
		return nil
	}
}

func (repository *OrderRepositoryMock) CreateOrder(order *Order) error {
	arguments := repository.Mock.Called(order)
	if arguments.Get(0) == nil {
		return errors.New("error CreateOrder")
	} else {
		return nil
	}
}

func (repository *OrderRepositoryMock) GetByOrderID(id string) (resultOrder []Order, err error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("error GetByOrderID")
	} else {
		order := arguments.Get(0).([]Order)
		return order, nil
	}
}
func (repository *OrderRepositoryMock) GetAll() (orderResult []Order, err error) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("error GetAll")
	} else {
		order := arguments.Get(0).([]Order)
		return order, nil
	}
}
func (repository *OrderRepositoryMock) Upsert(order *Order) error {
	arguments := repository.Mock.Called(order)
	if arguments.Get(0) == nil {
		return errors.New("error Upsert")
	} else {
		return nil
	}
}
func (repository *OrderRepositoryMock) FindAllOrderDetailByOrderID(id string) (resultOrderDetail []OrderDetail, err error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("error FindAllOrderDetailByOrderID")
	} else {
		orderDetail := arguments.Get(0).([]OrderDetail)
		return orderDetail, nil
	}
}

func (repository *OrderRepositoryMock) GetAllCart() (resultCart []cart.Cart, err error) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("error GetAllCart")
	} else {
		cart := arguments.Get(0).([]cart.Cart)
		return cart, nil
	}
}

func (repository *OrderRepositoryMock) GetProductByID(id string) (resultProduct *products.Products, err error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("error GetProductByID")
	} else {
		product := arguments.Get(0).(*products.Products)
		return product, nil
	}
}

func (repository *OrderRepositoryMock) UpdateProduct(product *products.Products) error {
	arguments := repository.Mock.Called(product)
	if arguments.Get(0) == nil {
		return errors.New("error UpdateProduct")
	} else {
		return nil
	}
}

func (repository *OrderRepositoryMock) FindPromoByID(id string) (resultsPromo *products.Promo, err error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("error FindPromoByID")
	} else {
		product := arguments.Get(0).(*products.Promo)
		return product, nil
	}
}
