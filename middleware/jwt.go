package middleware

import (
	"gin-mall/pkg/e"
	"gin-mall/pkg/utils"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorAuthToken
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthTokenTimeout
			}
		}
		if code != e.Success {
			c.JSON(e.Success, serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
