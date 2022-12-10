package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListFavorites(c *gin.Context) {
	listFavoriteService := service.FavoriteService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&listFavoriteService); err == nil {
		res := listFavoriteService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func CreateFavorite(c *gin.Context) {
	createFavoriteService := service.FavoriteService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&createFavoriteService); err != nil {
		res := createFavoriteService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	deleteFavoriteService := service.FavoriteService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := deleteFavoriteService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
