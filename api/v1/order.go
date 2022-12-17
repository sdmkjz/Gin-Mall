package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createOrderService := service.OrderService{}
	if err := c.ShouldBindJSON(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func ListOrder(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	listOrderService := service.OrderService{}
	if err := c.ShouldBindJSON(&listOrderService); err == nil {
		res := listOrderService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func GetOrder(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	showOrderService := service.OrderService{}
	res := showOrderService.Show(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func DeleteOrder(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	deleteOrderService := service.OrderService{}
	res := deleteOrderService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
