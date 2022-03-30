package pkg

import "github.com/gin-gonic/gin"

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func Error(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, Response{
		Success:    false,
		StatusCode: httpCode,
		Message:    err.Error(),
	})
}

func Success(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, Response{
		Success:    true,
		StatusCode: httpCode,
		Data:       data,
		Message:    "success",
	})
}
