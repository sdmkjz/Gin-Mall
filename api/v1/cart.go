package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCart(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createCartService := service.CartService{}
	if err := c.ShouldBindJSON(&createCartService); err == nil {
		res := createCartService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func UpdateCart(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	updateCartService := service.CartService{}
	if err := c.ShouldBindJSON(&updateCartService); err == nil {
		res := updateCartService.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

}

func DeleteCart(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	deleteCartService := service.CartService{}
	res := deleteCartService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ListCart(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	deleteCartService := service.CartService{}
	res := deleteCartService.List(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}
