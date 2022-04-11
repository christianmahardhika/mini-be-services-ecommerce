package products

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productsRepository = &ProductRepositoryMock{Mock: mock.Mock{}}
var productsService = useCase{repo: productsRepository}

func TestProductsService_GetFail(t *testing.T) {

	productsRepository.Mock.On("FindByKeyword", "Pembalut").Return(nil)

	result, err := productsService.SearchProducts("Pembalut")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestProductsService_GetSuccess(t *testing.T) {
	product := []*Products{
		{
			ID:      uuid.New(),
			Name:    "Product 1",
			Price:   100000,
			PromoID: uuid.New(),
			Stock:   10,
		},
		{
			ID:      uuid.New(),
			Name:    "Product 2",
			Price:   100000,
			PromoID: uuid.New(),
			Stock:   10,
		},
	}

	productsRepository.Mock.On("FindByKeyword", "Product").Return(product)

	result, err := productsService.SearchProducts("Product")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product[0].ID, result[0].ID)
	assert.Equal(t, product[1].Name, result[1].Name)
}
