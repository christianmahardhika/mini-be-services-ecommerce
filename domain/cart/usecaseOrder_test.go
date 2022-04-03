package cart

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var cartRepository = &CartRepositoryMock{Mock: mock.Mock{}}
var cartService = useCase{repo: cartRepository}

func TestCartService_GetCartInfoSuccess(t *testing.T) {
	cart := []Cart{
		Cart{
			ID:        9,
			ProductID: 140,
			Quantity:  8,
		},
		Cart{
			ID:        10,
			ProductID: 137,
			Quantity:  8,
		},
	}
	cartRepository.Mock.On("GetAll").Return(cart)

	result, err := cartService.GetCartInfo()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, cart[0].ID, result[0].ID)
	assert.Equal(t, cart[1].ProductID, result[1].ProductID)
}

func TestCartService_DeleteItemFromCartSuccess(t *testing.T) {
	cartRepository.Mock.On("Delete", int64(10)).Return(nil)

	result, err := cartService.DeleteItemFromCart(int64(10))
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCartService_ResetCartSuccess(t *testing.T) {
	cartRepository.Mock.On("DeleteAll").Return(nil)

	result, err := cartService.ResetCart()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCartService_AddItemToCartSuccess(t *testing.T) {
	item := Cart{
		ID:        9,
		ProductID: 140,
		Quantity:  8,
	}
	cartRepository.Mock.On("UpdateInsert", &item).Return(item)

	result, err := cartService.AddItemToCart(item)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
