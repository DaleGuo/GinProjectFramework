package service

import (
	"GinProjectFramework/db"
	"github.com/gin-contrib/sessions"
)

//返回登录结果和信息
func SignIn(userName string, password string,session sessions.Session) (bool,string) {
	result:=db.Authen(userName,password)
	if !result{
		return false,"用户名或密码错误"
	}

	//将用户登录信息存入session
	session.Set("hasSignIn", "true")
	session.Set("role", db.QueryRole(userName))
	session.Save()

	return true,"登录成功"
}