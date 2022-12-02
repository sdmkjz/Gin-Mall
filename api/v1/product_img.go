package v1

import (
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProductImg(c *gin.Context) {
	var listCatousel service.ListProductImg
	res := listCatousel.List(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}
