package main

import (
	"os"

	"github.com/christianmahardhika/mini-be-services-ecommerce/config"
	"github.com/christianmahardhika/mini-be-services-ecommerce/server"
	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine
var appConfig config.AppConfig

func init() {
	appConfig = server.InitiateConfig()
	ginEngine = server.InitiatializeGinServer(appConfig)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = appConfig.Port
	}

	server.StartApplication(ginEngine, port)
}
