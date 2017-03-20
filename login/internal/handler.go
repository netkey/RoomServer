package internal

import (
	"reflect"
	"github.com/name5566/leaf/gate"
	"RoomServer/msg"
	"RoomServer/db"
	"github.com/name5566/leaf/log"
	"RoomServer/common"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.CreateDeskReq{}, handleCreateDesk)
	handleMsg(&msg.CheckInDeskReq{}, handleCheckInDesk)
}



func handleCreateDesk(args []interface{}) {
	// 收到的消息
	m := args[0].(*msg.CreateDeskReq)
	// 消息的发送者
	a := args[1].(gate.Agent)
	tmpRes := msg.CreateDeskRes{}
	// 输出收到的消息的内容
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		a.WriteMsg(&tmpRes,common.EC_USER_NOT_LOGIN.Code())
		return
	}
	UserData := a.UserData().(db.User)

	a.SetUserData(UserData)
	a.WriteMsg(&tmpRes,common.EC_NONE.Code())
	go db.UserInfoNotice(UserData.ID,a)
	return
}


func handleCheckInDesk(args []interface{}) {
	// 收到的消息
	m := args[0].(*msg.CheckInDeskReq)
	// 消息的发送者
	a := args[1].(gate.Agent)
	tmpRes := msg.CheckInDeskRes{}
	// 输出收到的消息的内容
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		a.WriteMsg(&tmpRes,common.EC_USER_NOT_LOGIN.Code())
		return
	}
	UserData := a.UserData().(db.User)

	db.UserInit(&UserData)
	a.SetUserData(UserData)
	a.WriteMsg(&tmpRes,common.EC_NONE.Code())
	go db.UserInfoNotice(UserData.ID,a)
	return
}
