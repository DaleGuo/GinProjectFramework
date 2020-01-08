package service

import (
	"GinProjectFramework/db"
	"github.com/gin-contrib/sessions"
)

//用户登录
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

//用户退出
func SignOut(session sessions.Session) (bool,string) {
	//将用户登录信息从session删除
	session.Set("hasSignIn", "")
	session.Set("role", "")
	session.Save()

	return true,"退出成功"
}