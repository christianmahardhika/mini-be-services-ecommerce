package order

import (
	"testing"

	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/cart"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var orderRepository = &OrderRepositoryMock{Mock: mock.Mock{}}
var orderService = useCase{repo: orderRepository}

func TestOrderService_PlaceOrderSuccess(t *testing.T) {
	order := Order{
		UserID:    "1",
		TotalCart: 0,
	}
	cart := []cart.Cart{
		{
			ID:        uuid.New().String(),
			ProductID: uuid.New().String(),
			Quantity:  8,
		},
		{
			ID:        uuid.New().String(),
			ProductID: uuid.New().String(),
			Quantity:  8,
		},
	}

	product1 := products.Products{
		ID:      cart[0].ProductID,
		PromoID: uuid.New().String(),
		Name:    "Product 1",
		Price:   100000,
		Stock:   10,
	}
	product2 := products.Products{
		ID:      cart[1].ProductID,
		PromoID: uuid.New().String(),
		Name:    "Product 2",
		Price:   100000,
		Stock:   10,
	}
	productUpdated1 := products.Products{
		ID:      product1.ID,
		PromoID: product1.PromoID,
		Name:    "Product 1",
		Price:   100000,
		Stock:   product1.Stock - cart[0].Quantity,
	}
	productUpdated2 := products.Products{
		ID:      product2.ID,
		PromoID: product2.PromoID,
		Name:    "Product 2",
		Price:   100000,
		Stock:   product2.Stock - cart[1].Quantity,
	}

	promo1 := products.Promo{
		ID:        product1.PromoID,
		Discount:  10,
		PromoCode: "PROMO1",
		PromoType: "percentage",
	}
	promo2 := products.Promo{
		ID:        product2.PromoID,
		Discount:  20,
		PromoCode: "PROMO2",
		PromoType: "percentage",
	}

	orderDetail := []OrderDetail{
		{
			OrderID:   order.ID,
			ProductID: cart[0].ProductID,
			Quantity:  8,
			Price:     product1.Price,
			PromoCode: promo1.PromoCode,
			Total:     product1.Price - (product1.Price * (promo1.Discount / 100)),
		},
		{
			OrderID:   order.ID,
			ProductID: cart[1].ProductID,
			Quantity:  8,
			Price:     product2.Price,
			PromoCode: promo2.PromoCode,
			Total:     product2.Price - (product2.Price * (promo2.Discount / 100)),
		},
	}
	orderDetail1 := OrderDetail{
		// OrderID:   order.ID,
		ProductID: cart[0].ProductID,
		Quantity:  cart[0].Quantity,
		Price:     product1.Price,
		PromoCode: promo1.PromoCode,
		Total:     product1.Price - (product1.Price * (promo1.Discount / 100)),
	}
	orderDetail2 := OrderDetail{

		// OrderID:   order.ID,
		ProductID: cart[1].ProductID,
		Quantity:  cart[1].Quantity,
		Price:     product2.Price,
		PromoCode: promo2.PromoCode,
		Total:     product2.Price - (product2.Price * (promo2.Discount / 100)),
	}
	orderTotal := Order{
		UserID:    order.UserID,
		TotalCart: orderDetail[0].Total + orderDetail[1].Total,
	}

	InvoiceOrder := InvoiceOrder{
		Order:       orderTotal,
		OrderDetail: orderDetail,
	}

	orderRepository.Mock.On("CreateOrder", &order).Return(order)
	orderRepository.Mock.On("GetAllCart").Return(cart)
	orderRepository.Mock.On("GetProductByID", product1.ID).Return(&product1)
	orderRepository.Mock.On("GetProductByID", product2.ID).Return(&product2)
	orderRepository.Mock.On("UpdateProduct", &productUpdated1).Return(nil)
	orderRepository.Mock.On("UpdateProduct", &productUpdated2).Return(nil)
	orderRepository.Mock.On("FindPromoByID", product1.PromoID).Return(&promo1)
	orderRepository.Mock.On("FindPromoByID", product2.PromoID).Return(&promo2)
	orderRepository.Mock.On("CreateOrderDetail", &orderDetail1).Return(orderDetail1)
	orderRepository.Mock.On("CreateOrderDetail", &orderDetail2).Return(orderDetail2)
	orderRepository.Mock.On("Upsert", &orderTotal).Return(orderTotal)
	orderRepository.Mock.On("FindAllOrderDetailByOrderID", order.ID).Return(orderDetail, nil)

	result, err := orderService.PlaceOrder()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &InvoiceOrder, result)
}
