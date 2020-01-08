package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"strings"
)

//鉴权
func AuthzMiddleWare(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//登录功能跳过验证
		if url := c.Request.URL.String(); strings.HasPrefix(url,"/ginFrameWork/login") || strings.HasPrefix(url,"/ginFrameWork/signIn"){
			c.Next()
			return
		}

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := sessions.Default(c).Get("role")

		//静态资源文件跳过验证
		if strings.HasPrefix(obj,"/static"){
			c.Next()
			return
		}

		//判断策略中是否存在
		if e.Enforce(sub, obj, act) {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{"status":1,"message": "用户无权限"})
			c.Abort()
		}
	}
}