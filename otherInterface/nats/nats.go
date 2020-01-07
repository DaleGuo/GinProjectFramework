package nats

import (
	"log"
	"github.com/nats-io/go-nats"
	"encoding/json"
	"GinProjectFramework/global"
)

var natsConn *nats.Conn    //nats连接

func InitNats() error {
	natsURL, err := global.GetNatsURL()
	if err != nil {
		log.Println("nats连接失败", err)
		return err
	}

	nc, err := nats.Connect(natsURL)
	if nc == nil {
		log.Println("nats连接失败", err)
		return err
	}

	natsConn = nc
	log.Println("nats连接成功")

	subscribe()//订阅

	return nil
}

func subscribe() {
	natsTopic,err:=global.GetNatsTopic()
	if err!=nil{
		log.Println("获取nats订阅话题失败", err)
		return
	}
	natsConn.Subscribe(natsTopic, natsCallback)

	//todo:其它订阅
}

//订阅回调函数
func natsCallback(m *nats.Msg) {
	var data Cmd
	err:=json.Unmarshal(m.Data,&data)
	if err!=nil{
		log.Println("参数解析失败",err)
	}else{
		//todo:处理命令
	}
}

//todo：其它回调函数