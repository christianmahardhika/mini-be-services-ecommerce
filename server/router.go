package server

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/cart"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/order"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"github.com/gin-gonic/gin"
)

func initiateRouter(r *gin.Engine, productController products.Controller, cartController cart.Controller, orderController order.Controller) *gin.Engine {

	// Product Domain Routes
	r.GET("/products", productController.SearchProducts)

	// Cart Domain Routes
	r.POST("/cart/add", cartController.AddItemToCart)
	r.GET("/cart/reset", cartController.ResetCart)
	r.GET("/cart/delete", cartController.DeleteItemFromCart)
	r.GET("/cart", cartController.GetCartInfo)

	// Order Domain Routes
	r.GET("/order/checkout", orderController.CreateOrder)

	return r
}
