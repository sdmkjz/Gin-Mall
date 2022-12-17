package routes

import (
	api "gin-mall/api/v1"
	"gin-mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		v1.GET("carousels", api.ListCarousel)
		// 商品操作
		v1.GET("products", api.ListProduct)
		v1.GET("product/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)
			authed.POST("money", api.ShowMoney)
			// 商品操作
			authed.POST("product", api.CreateProduct)
			authed.POST("products", api.SearchProduct)
			// 收藏夹操作
			authed.GET("favorites", api.ListFavorites)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)
			// 地址簿操作
			authed.POST("address", api.CreateAddress)
			authed.GET("address", api.ListAddress)
			authed.GET("address/:id", api.GetAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.DELETE("address/:id", api.DeleteAddress)
			// 购物车操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts", api.ListCart)
			authed.PUT("carts/:id", api.UpdateCart)
			authed.DELETE("carts/:id", api.DeleteCart)
			// 订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders", api.ListOrder)
			authed.GET("orders/:id", api.GetOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)
			// 支付
			authed.POST("paydown", api.OrderPay)
		}
	}
	return r
}
