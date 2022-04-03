package order

import (
	"net/http"

	"github.com/christianmahardhika/mini-be-services-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UseCase UseCase
}

func (ctrl *Controller) CreateOrder(c *gin.Context) {

	result, err := ctrl.UseCase.PlaceOrder()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, http.StatusCreated, result)
}
