package products

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindByKeyword(keyword string) (products []*Products, err error) {
	arguments := repository.Mock.Called(keyword)
	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	} else {
		products := arguments.Get(0).([]*Products)
		return products, nil
	}
}
