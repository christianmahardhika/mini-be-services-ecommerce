package helloworld

import "github.com/gin-gonic/gin"

type Controller struct {
}

func (ctr Controller) HeloFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world.",
	})
}
