package main

import (
	"GinProjectFramework/router"
	"GinProjectFramework/otherInterface/websocket"
	"GinProjectFramework/otherInterface/nats"
	"GinProjectFramework/db"
)

func main() {
	nats.InitNats()//nats

	websocket.InitWS()//websocket，新启动了一个goroutine

	db.InitDB()//mysql

	router.Route()//gin
}