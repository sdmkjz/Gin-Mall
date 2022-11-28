package v1

import (
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService
	res := listCarousel.List(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
