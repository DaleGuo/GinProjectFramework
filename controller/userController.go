package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"GinProjectFramework/service"
	"github.com/gin-contrib/sessions"
)

//登录
func SignIn(c *gin.Context)  {
	userName:=c.PostForm("userName")
	password:=c.PostForm("password")

	result,message:=service.SignIn(userName,password,sessions.Default(c))
	if result {
		c.JSON(http.StatusOK,gin.H{})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"message":message,
		})
	}
}