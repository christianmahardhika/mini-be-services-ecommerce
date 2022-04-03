package cart

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type CartRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CartRepositoryMock) GetAll() ([]Cart, error) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	} else {
		cart := arguments.Get(0).([]Cart)
		return cart, nil
	}
}

func (repository *CartRepositoryMock) GetByID(cartID string) (*Cart, error) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	} else {
		cart := arguments.Get(0).(Cart)
		return &cart, nil
	}
}
func (repository *CartRepositoryMock) UpdateInsert(cart *Cart) error {
	arguments := repository.Mock.Called(cart)
	if arguments.Get(0) == nil {
		return errors.New("Error")
	} else {
		return nil
	}
}
func (repository *CartRepositoryMock) Delete(cartID int64) error {
	arguments := repository.Mock.Called(cartID)
	if arguments.Get(0) == nil {
		return nil
	} else {
		return errors.New("Error")
	}
}
func (repository *CartRepositoryMock) DeleteAll() error {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	} else {
		return errors.New("Error")
	}
}
