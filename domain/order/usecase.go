package order

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/christianmahardhika/mini-be-services-ecommerce/pkg/helper"
	"gorm.io/gorm"
)

type UseCase interface {
	PlaceOrder() (resultOrder *[]Order, err error)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo}
}

type useCase struct {
	repo Repository
}

// PlaceOrder implements UseCase
func (uc *useCase) PlaceOrder() (*[]Order, error) {

	// Get latest order ID
	var orderID string
	resOrderId, err := uc.repo.GetLatestOrderID()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			orderID = "ODR0001"
		} else {
			return nil, errors.New("failed to get latest order ID")
		}
	} else {
		orderID = resOrderId.OrderID
		latestOrderID := orderID[len(orderID)-1:]
		latestOrderIDInt, err := strconv.Atoi(latestOrderID)
		if err != nil {
			return nil, err
		}
		orderID = "ODR000" + strconv.Itoa(int(latestOrderIDInt)+1)
	}
	// Create order

	res, err := uc.repo.GetAllCart()
	if err != nil {
		return nil, errors.New("failed to get cart")
	}

	var totalCart int64

	for _, cart := range res {
		var order Order
		resProduct, err := uc.repo.GetProductByID(cart.ProductID)
		if err != nil {
			return nil, err
		}
		//check stock
		if resProduct.Stock < cart.Quantity {
			return nil, errors.New("stock is not enough")
		}
		resProduct.Stock -= cart.Quantity
		uc.repo.UpdateProduct(&resProduct)

		// calculate Promo Price

		totalResultPrice, err := helper.PromoHelper(resProduct.Promo, resProduct.Price, cart.Quantity)
		if err != nil {
			return nil, err
		}

		order.OrderID = orderID
		order.ProductID = cart.ProductID
		order.Quantity = cart.Quantity
		order.Price = resProduct.Price
		order.PromoCode = resProduct.Promo
		order.Total = totalResultPrice

		totalCart = totalCart + totalResultPrice

		err = uc.repo.Create(&order)
		if err != nil {
			return nil, err
		}
	}

	// calculate total cart
	fmt.Println(totalCart)
	resultOrder, err := uc.repo.GetByOrderID(orderID)
	for _, order := range resultOrder {
		order.TotalCart = totalCart
		err = uc.repo.Upsert(&order)
	}

	// show all order
	resultOrder, err = uc.repo.GetByOrderID(orderID)
	return &resultOrder, err

}
