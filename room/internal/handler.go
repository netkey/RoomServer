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
	handleMsg(&msg.LoginReq{}, handleLogin)
}



func handleLogin(args []interface{}) {
	// 收到的消息
	m := args[0].(*msg.LoginReq)
	// 消息的发送者
	a := args[1].(gate.Agent)
	tmpRes := msg.LoginRes{}
	// 输出收到的消息的内容
	UserData := db.User{}

	if m.LoginType == int32(msg.LoginTypes_Face_Book){
		//facebook login
		log.Debug("FaceBook Login Name :%s  Token:%s", m.UserName,m.AccessToken)
		tmpUserData,err := db.GetPlayerByUserName(m.UserName)
		if err != nil{
			log.Debug("GetUser error")
			a.WriteMsg(&tmpRes,common.EC_DB.Code())
			return
		}
		UserData = tmpUserData
	}else if m.LoginType == int32(msg.LoginTypes_Normal_Login){
		log.Debug("Name :%s  Pass:%s", m.UserName,m.PassWd)
		tmpUserData,err := db.GetPlayerByUserName(m.UserName)
		if err != nil{
			log.Debug("GetUser error")
			a.WriteMsg(&tmpRes,common.EC_DB.Code())
			return
		}
		if m.PassWd != tmpUserData.PassWd{
			log.Debug("GetUser EC_PW_NOT_MATCH")
			a.WriteMsg(&tmpRes,common.EC_PW_NOT_MATCH.Code())
			return
		}
		UserData = tmpUserData
	}else if m.LoginType == int32(msg.LoginTypes_ReConnect){
		log.Debug("Reconnect Login Name :%s  Token:%s", m.UserName,m.ReconnectToken)
		tmpUserData,err := db.GetPlayerByUserName(m.UserName)
		if err != nil{
			log.Debug("GetUser error")
			a.WriteMsg(&tmpRes,common.EC_DB.Code())
			return
		}
		UserData = tmpUserData
	}
	log.Debug("GetUser: %v",UserData)
	log.Debug("GetUser EC_NONE")
	db.UserInit(&UserData)
	a.SetUserData(UserData)
	a.WriteMsg(&tmpRes,common.EC_NONE.Code())
	go db.UserInfoNotice(UserData.ID,a)
	return
}

