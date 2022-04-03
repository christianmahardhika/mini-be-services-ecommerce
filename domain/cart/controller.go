package cart

import (
	"net/http"
	"strconv"

	"github.com/christianmahardhika/mini-be-services-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UseCase UseCase
}

func (ctrl *Controller) AddItemToCart(c *gin.Context) {
	var itemCart Cart
	err := c.ShouldBindJSON(&itemCart)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	result, err := ctrl.UseCase.AddItemToCart(itemCart)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, http.StatusCreated, result)
}

func (ctrl *Controller) DeleteItemFromCart(c *gin.Context) {
	itemID, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	result, err := ctrl.UseCase.DeleteItemFromCart(itemID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, http.StatusCreated, result)
}

func (ctrl *Controller) GetCartInfo(c *gin.Context) {
	result, err := ctrl.UseCase.GetCartInfo()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, http.StatusOK, result)
}

func (ctrl *Controller) ResetCart(c *gin.Context) {
	result, err := ctrl.UseCase.ResetCart()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, http.StatusCreated, result)
}
