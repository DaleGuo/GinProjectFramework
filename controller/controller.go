package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHtml(c *gin.Context)  {
	c.HTML(http.StatusOK,"login.html",gin.H{})
}

func IndexHtml(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",gin.H{})
}

//todo:其它页面发布回调函数