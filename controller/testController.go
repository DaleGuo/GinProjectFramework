package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"GinProjectFramework/service"
)

func GetResource1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "请求成功",
	})
}
func PostResource1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "请求成功",
	})
}
func GetResource2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "请求成功",
	})
}
func PostResource2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "请求成功",
	})
}

func EnablePostResource1(c *gin.Context) {
	_,message:=service.AddPolicy("student","/ginFrameWork/resource1","POST")
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": message,
	})
}

func DisablePostResource1(c *gin.Context) {
	_,message:=service.RemovePolicy("student","/ginFrameWork/resource1","POST")
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": message,
	})
}