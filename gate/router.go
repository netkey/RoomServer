package gate

import (
	"RoomServer/game"
	"RoomServer/msg"
	"RoomServer/login"
	"RoomServer/room"
)

func init() {
	// 这里指定消息 Room 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.LoginReq{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.CreateDeskReq{}, room.ChanRPC)
	msg.Processor.SetRouter(&msg.CheckInDeskReq{}, room.ChanRPC)
	msg.Processor.SetRouter(&msg.GetVipShopListReq{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeartBeat{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.GMGameReq{},game.ChanRPC)
	//msg.Processor.SetRouter(&msg.MarryGameReq{},game.ChanRPC)
}