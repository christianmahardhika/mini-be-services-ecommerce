package products

import (
	"net/http"

	"github.com/christianmahardhika/mini-be-services-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UseCase UseCase
}

func (ctr Controller) SearchProducts(c *gin.Context) {
	keyword := c.Query("keyword")
	results, err := ctr.UseCase.SearchProducts(keyword)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, http.StatusOK, results)
}
