package server

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/config"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/cart"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/order"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitiatializeGinServer(appConfig config.AppConfig) *gin.Engine {
	// InitiateGinServer is a function that initiates the gin server.

	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = appConfig.AllowedOrigins
	ginConfig.AllowHeaders = appConfig.AllowedHeaders

	r := gin.Default()
	r.Use(cors.New(ginConfig))
	dbConnection := GetDbConnection(appConfig.DBConfig)

	// Instantiate the usecase
	productUC := products.NewUseCase(products.NewRepository(dbConnection))
	cartUC := cart.NewUseCase(cart.NewRepository(dbConnection))
	orderUC := order.NewUseCase(order.NewRepository(dbConnection))

	// Instantiate the controller
	productController := products.Controller{UseCase: productUC}
	cartController := cart.Controller{UseCase: cartUC}
	orderController := order.Controller{UseCase: orderUC}

	// call the router
	return initiateRouter(r, productController, cartController, orderController)
}
