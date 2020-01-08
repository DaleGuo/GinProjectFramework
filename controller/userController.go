package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"GinProjectFramework/service"
	"github.com/gin-contrib/sessions"
	"io/ioutil"
	"encoding/json"
)

type signInStruct struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

//登录
func SignIn(c *gin.Context)  {
	s, _ := ioutil.ReadAll(c.Request.Body)

	var data signInStruct
	err:=json.Unmarshal(s,&data)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"status":1,
			"message":"请求参数格式非法",
		})
		return
	}

	result,message:=service.SignIn(data.UserName,data.Password,sessions.Default(c))
	if result {
		c.JSON(http.StatusOK,gin.H{
			"status":0,
			"message":message,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":1,
			"message":message,
		})
	}
}