package websocket

import (
	"net/http"
	"github.com/gorilla/websocket"
	"encoding/json"
	"log"
	"GinProjectFramework/global"
)

var wsConn *websocket.Conn//webSocket连接

// 解决跨域问题
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket轮询处理函数
func wsHandle(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil{
			log.Println(err)
			c.Close()
			return
		}

		var messageJson map[string] interface{}
		err = json.Unmarshal(message,&messageJson)
		if messageJson == nil {
			continue
		}
		cmdID := messageJson["cmdID"].(float64)
		switch cmdID {
		case ConnectedCmd:
			wsConn = c
			log.Println("websocket连接成功")
		}
		//todo：其它命令处理
	}
}

//初始化webSocket
func InitWS() error {
	wsUrl,err:=global.GetWebSocketURL()
	if err!=nil{
		log.Println("获取webSocket地址失败",err)
		return err
	}
	http.HandleFunc(wsUrl, wsHandle)

	wsPort,err:=global.GetWebSocketPort()
	if err!=nil{
		log.Println("获取webSocket端口号失败",err)
		return err
	}
	go http.ListenAndServe(wsPort,nil)

	log.Println("webSocket服务启动成功")
	return nil
}

//通过webSocket发送数据，data需定义json标注
func WSSendJson(data interface{}) error {
	if wsConn==nil{
		return ErrNullPointer
	}
	return wsConn.WriteJSON(data)
}