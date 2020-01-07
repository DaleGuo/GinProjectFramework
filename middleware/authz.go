package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
)

//鉴权
func AuthzMiddleWare(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := sessions.Default(c).Get("role")

		//判断策略中是否存在
		if e.Enforce(sub, obj, act) {
			c.Next()
		} else {
			c.Abort()
		}
	}
}