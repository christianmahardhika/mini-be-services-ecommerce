package server

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/config"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/helloworld"
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
	GetDbConnection(appConfig.DBConfig)

	helloworldController := helloworld.Controller{}
	return initiateRouter(r, helloworldController)
}
