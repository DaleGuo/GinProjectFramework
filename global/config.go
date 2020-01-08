package global

import (
	"github.com/go-ini/ini"
	"fmt"
)

//获取nats服务器地址
func GetNatsURL() (string,error) {
	cfg,err:=ini.Load("./config/serverConfig.ini")
	if err!=nil{
		return "",err
	}
	natsUrl := cfg.Section("Nats").Key("url").String()
	return natsUrl,nil
}

//获取nats订阅话题名称
func GetNatsTopic() (string,error) {
	cfg,err:=ini.Load("./config/serverConfig.ini")
	if err!=nil{
		return "",err
	}
	natsUrl := cfg.Section("Nats").Key("topic").String()
	return natsUrl,nil
}

//获取http路由端口号
func GetWebRoutePort() (string,error) {
	cfg,err:=ini.Load("./config/serverConfig.ini")
	if err!=nil{
		return "",err
	}
	webRoutePort := cfg.Section("webRoute").Key("port").String()
	return webRoutePort,nil
}

//获取webSocket地址
func GetWebSocketURL() (string,error) {
	cfg,err:=ini.Load("./config/serverConfig.ini")
	if err!=nil{
		return "",err
	}
	wsUrl := cfg.Section("webSocket").Key("url").String()
	return wsUrl,nil
}

//获取webSocket端口号
func GetWebSocketPort() (string,error) {
	cfg,err:=ini.Load("./config/serverConfig.ini")
	if err!=nil{
		return "",err
	}
	wsPort := cfg.Section("webSocket").Key("port").String()
	return wsPort,nil
}

//获取数据库连接Url
func GetDBConnectUrl() (string,error) {
	cfg,err:=ini.Load("./config/serverConfig.ini")
	if err!=nil{
		return "",err
	}
	section:=cfg.Section("mysql")

	user := section.Key("user").String()
	password := section.Key("password").String()
	address := section.Key("address").String()
	tableName := section.Key("dbName").String()

	url:=fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",user,password,address,tableName)
	return url,nil
}

//获取数据库连接Url
func GetAccessPolicyUrl() (string,error) {
	cfg,err:=ini.Load("./config/serverConfig.ini")
	if err!=nil{
		return "",err
	}
	section:=cfg.Section("mysql")

	user := section.Key("user").String()
	password := section.Key("password").String()
	address := section.Key("address").String()
	tableName := section.Key("tableName").String()

	url:=fmt.Sprintf("%s:%s@tcp(%s)/%s",user,password,address,tableName)
	return url,nil
}