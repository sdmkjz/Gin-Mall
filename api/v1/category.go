package v1

import (
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(c *gin.Context) {
	var listCategory service.CategoryService
	res := listCategory.List(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
