package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"RoomServer/msg"
	"RoomServer/common"
)

func init() {
	// 向当前模块（game 模块）注册 Room 消息的消息处理函数 handleRoom
	handler(&msg.HeartBeat{},handleHeartBeat)
	handler(&msg.GMGameReq{},handleGmCmd)

}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHeartBeat(args []interface{}){
	//m := args[0].(*msg.HeartBeat)
	a := args[1].(gate.Agent)
	tmpRes := msg.HeartBeat{}
	a.WriteMsg(&tmpRes,common.EC_NONE.Code())
	return
}


