package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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