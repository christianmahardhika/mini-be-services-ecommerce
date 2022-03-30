package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

func StartApplication(ginEng *gin.Engine, port string) {
	ginEngine = ginEng
	address := fmt.Sprintf(":%v", port)
	err := ginEngine.Run(address)
	if err != nil {
		panic(err)
	}
}
