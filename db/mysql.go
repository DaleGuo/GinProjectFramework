package db

import (
	"github.com/jinzhu/gorm"
	"GinProjectFramework/global"
	"log"
)

var db *gorm.DB

func InitDB()(error) {
	dbUrl,err:=global.GetDBConnectUrl()
	if err!=nil{
		log.Println("读取数据库连接url失败",err)
		return err
	}

	db,err = gorm.Open("mysql",dbUrl)
	if err!=nil{
		log.Println("数据库连接失败",err)
		return err
	}

	log.Println("数据库连接成功")
	return nil
}