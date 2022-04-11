package order

import (
	"errors"

	"github.com/christianmahardhika/mini-be-services-ecommerce/pkg/helper"
)

type UseCase interface {
	PlaceOrder() (resultOrder *InvoiceOrder, err error)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo}
}

type useCase struct {
	repo Repository
}

// PlaceOrder implements UseCase
func (uc *useCase) PlaceOrder() (*InvoiceOrder, error) {

	// Create order
	var order Order
	UserId := "1"
	order = Order{
		UserID:    UserId,
		TotalCart: 0,
	}

	// uc.repo.CreateOrder(&order)

	err := uc.repo.CreateOrder(&order)
	if err != nil {
		return nil, err
	}

	res, err := uc.repo.GetAllCart()
	if err != nil {
		return nil, errors.New("failed to get cart")
	}

	var totalCart int64

	for _, cart := range res {
		var orderDetail OrderDetail
		resProduct, err := uc.repo.GetProductByID(cart.ProductID)
		if err != nil {
			return nil, err
		}
		//check stock
		if resProduct.Stock < cart.Quantity {
			return nil, errors.New("stock is not enough")
		}
		resProduct.Stock -= cart.Quantity
		uc.repo.UpdateProduct(resProduct)

		// find promo
		resPromo, err := uc.repo.FindPromoByID(resProduct.PromoID)
		if err != nil {
			return nil, err
		}
		// calculate Promo Price

		totalResultPrice, err := helper.PromoHelper(resPromo.Discount, resProduct.Price, cart.Quantity)
		if err != nil {
			return nil, err
		}
		orderDetail.OrderID = order.ID
		orderDetail.ProductID = cart.ProductID
		orderDetail.Quantity = cart.Quantity
		orderDetail.Price = resProduct.Price
		orderDetail.PromoCode = resPromo.PromoCode
		orderDetail.Total = totalResultPrice

		totalCart = totalCart + totalResultPrice

		// uc.repo.CreateOrderDetail(&orderDetail)

		err = uc.repo.CreateOrderDetail(&orderDetail)
		if err != nil {
			return nil, err
		}
	}

	// update total cart
	order.TotalCart = totalCart
	uc.repo.Upsert(&order)

	// show all order
	resOrderDetail, err := uc.repo.FindAllOrderDetailByOrderID(order.ID)
	if err != nil {
		return nil, err
	}

	var InvoiceOrder InvoiceOrder

	InvoiceOrder.Order = order
	InvoiceOrder.OrderDetail = resOrderDetail
	return &InvoiceOrder, nil

}
