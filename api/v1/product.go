package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func ListProduct(c *gin.Context) {
	ListProductService := service.ProductService{}
	if err := c.ShouldBindJSON(&ListProductService); err == nil {
		res := ListProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func SearchProduct(c *gin.Context) {
	SearchProductService := service.ProductService{}
	if err := c.ShouldBindJSON(&SearchProductService); err == nil {
		res := SearchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func ShowProduct(c *gin.Context) {
	ShowProductService := service.ProductService{}
	res := ShowProductService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}
