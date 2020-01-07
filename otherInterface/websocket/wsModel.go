package websocket

import "errors"

//webSocket cmd
type Cmd struct {
	CMD float64 `json:"cmdID"`
	Data interface{} `json:"data"`
}

//定义webSocket命令
const (
	ConnectedCmd  = iota
	//todo:其它命令定义
)

//webSocket errors
var (
	ErrNullPointer = errors.New("NullPointer:webSocket has disconnected")
	//todo:其它webSocket错误定义
)