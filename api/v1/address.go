package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	if err := c.ShouldBindJSON(&createAddressService); err == nil {
		res := createAddressService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func GetAddress(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	showAddressService := service.AddressService{}
	res := showAddressService.Show(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ListAddress(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	listAddressService := service.AddressService{}
	res := listAddressService.List(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}

func UpdateAddress(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	updateAddressService := service.AddressService{}
	res := updateAddressService.Update(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func DeleteAddress(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	deleteAddressService := service.AddressService{}
	res := deleteAddressService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
