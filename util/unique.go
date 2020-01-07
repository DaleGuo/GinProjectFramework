package util

import "github.com/satori/go.uuid"

//创建唯一标识符
func CreateUniqueID() (string,error) {
	sessionID,err:=uuid.NewV4()
	return sessionID.String(),err
}