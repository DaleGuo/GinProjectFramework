package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"
)

//认证
func AuthenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session := sessions.Default(c); session.Get("hasSignIn") == "true" {
			c.Next()
			return
		}

		if url := c.Request.URL.String(); url == "/ginFrameWork/signIn" || url == "/ginFrameWork/login"{
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户未登录"})
		c.Abort()
		return
	}
}