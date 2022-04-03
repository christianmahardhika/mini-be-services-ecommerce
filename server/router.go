package server

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/cart"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/helloworld"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"github.com/gin-gonic/gin"
)

func initiateRouter(r *gin.Engine, helloworldController helloworld.Controller, productController products.Controller, cartController cart.Controller) *gin.Engine {
	r.GET("/", helloworldController.HeloFunction)

	// Product Domain Routes
	r.GET("/products", productController.SearchProducts)

	// Cart Domain Routes
	r.POST("/cart/add", cartController.AddItemToCart)
	r.GET("/cart/reset", cartController.ResetCart)
	r.GET("/cart/delete", cartController.DeleteItemFromCart)
	r.GET("/cart", cartController.GetCartInfo)

	return r
}
