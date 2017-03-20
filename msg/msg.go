package msg

import (
	"github.com/name5566/leaf/network/protobuf"

)

// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&LoginReq{},uint16(MsgID_loginReq))
	Processor.Register(&LoginRes{},uint16(MsgID_loginRes))
	Processor.Register(&PlayerInfoUpdate{},uint16(MsgID_playerInfoUpdate))
	Processor.Register(&GetVipShopListReq{},uint16(MsgID_getVipShopListReq))
	Processor.Register(&GetVipShopListRes{},uint16(MsgID_getVipShopListRes))
	Processor.Register(&CreateDeskReq{},uint16(MsgID_createDeskReq))
	Processor.Register(&CreateDeskRes{},uint16(MsgID_createDeskRes))
	Processor.Register(&CheckInDeskReq{},uint16(MsgID_checkInDeskReq))
	Processor.Register(&CheckInDeskRes{},uint16(MsgID_checkInDeskRes))
	Processor.Register(&HeartBeat{},uint16(MsgID_heartBeat))
	Processor.Register(&GMGameReq{},uint16(MsgID_gmCmdReq))
	Processor.Register(&GMGameRes{},uint16(MsgID_gmCmdRes))
}