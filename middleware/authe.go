package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"
	"strings"
)

//认证
func AuthenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session := sessions.Default(c); session.Get("hasSignIn") == "true" {
			c.Next()
			return
		}

		//登录页面和登录、退出操作跳过验证
		if url := c.Request.URL.String(); strings.HasPrefix(url, "/ginFrameWork/login") || strings.HasPrefix(url, "/ginFrameWork/signIn") || strings.HasPrefix(url, "/ginFrameWork/signOut") {
			c.Next()
			return
		}

		//静态资源文件跳过验证
		url := c.Request.URL.RequestURI()
		if strings.HasPrefix(url,"/static"){
			c.Next()
			return
		}

		c.HTML(http.StatusOK,"login.html",gin.H{"status":1,"message": "用户未登录"})
		c.Abort()
		return
	}
}