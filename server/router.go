package server

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/helloworld"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"github.com/gin-gonic/gin"
)

func initiateRouter(r *gin.Engine, helloworldController helloworld.Controller, productController products.Controller) *gin.Engine {
	r.GET("/", helloworldController.HeloFunction)

	r.GET("/products", productController.SearchProducts)

	return r
}
