package server

import (
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/helloworld"
	"github.com/gin-gonic/gin"
)

func initiateRouter(r *gin.Engine, helloworldController helloworld.Controller) *gin.Engine {
	r.GET("/", helloworldController.HeloFunction)
	return r
}
