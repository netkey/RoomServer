package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"RoomServer/msg"
)

func init() {
	// 向当前模块（game 模块）注册 Room 消息的消息处理函数 handleRoom
	handler(&msg.LoginReq{}, handleLogin)
	handler(&msg.OneArmPlayReq{},handlePlay)
	handler(&msg.HeartBeat{},handleHeartBeat)
	handler(&msg.GMGameReq{},handleGmCmd)
	//handler(&msg.MarryGameReq{},handleMarryGame)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHeartBeat(args []interface{}){
	//m := args[0].(*msg.HeartBeat)
	a := args[1].(gate.Agent)
	tmpRes := msg.HeartBeat{}
	a.WriteMsg(&tmpRes,EC_NONE.Code())
	return
}


func handleLogin(args []interface{}) {
	// 收到的消息
	m := args[0].(*msg.LoginReq)
	// 消息的发送者
	a := args[1].(gate.Agent)
	tmpRes := msg.LoginRes{}
	// 输出收到的消息的内容
	UserData := User{}

	if m.LoginType == int32(msg.LoginTypes_Face_Book){
		//facebook login
		log.Debug("FaceBook Login Name :%s  Token:%s", m.UserName,m.AccessToken)
		tmpUserData,err := GetPlayerByUserName(m.UserName)
		if err != nil{
			log.Debug("GetUser error")
			a.WriteMsg(&tmpRes,EC_DB.Code())
			return
		}
		UserData = tmpUserData
	}else if m.LoginType == int32(msg.LoginTypes_Normal_Login){
		log.Debug("Name :%s  Pass:%s", m.UserName,m.PassWd)
		tmpUserData,err := GetPlayerByUserName(m.UserName)
		if err != nil{
			log.Debug("GetUser error")
			a.WriteMsg(&tmpRes,EC_DB.Code())
			return
		}
		if m.PassWd != tmpUserData.PassWd{
			log.Debug("GetUser EC_PW_NOT_MATCH")
			a.WriteMsg(&tmpRes,EC_PW_NOT_MATCH.Code())
			return
		}
		UserData = tmpUserData
	}else if m.LoginType == int32(msg.LoginTypes_ReConnect){
		log.Debug("Reconnect Login Name :%s  Token:%s", m.UserName,m.ReconnectToken)
		tmpUserData,err := GetPlayerByUserName(m.UserName)
		if err != nil{
			log.Debug("GetUser error")
			a.WriteMsg(&tmpRes,EC_DB.Code())
			return
		}
		UserData = tmpUserData
	}
	log.Debug("GetUser: %v",UserData)
	log.Debug("GetUser EC_NONE")
	UserInit(&UserData)
	a.SetUserData(UserData)
	a.WriteMsg(&tmpRes,EC_NONE.Code())
	go UserInfoNotice(UserData.ID,a)
	return
}


func handlePlay(args []interface{}){
	// 收到的消息
	m := args[0].(*msg.OneArmPlayReq)
	// 消息的发送者
	a := args[1].(gate.Agent)
	log.Debug("LineCnt :%d  Betting:%d", m.LineCnt,m.Betting)
	tmpRes := msg.OneArmPlayRes{
		ShowTable:&msg.PlayTableData{},
		ResultDat:&msg.ResultData{},
	}
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		a.WriteMsg(&tmpRes,EC_USER_NOT_LOGIN.Code())
		return
	}
	tmpData := a.UserData().(User)
	if  tmpData.ID == ``{
		log.Debug(`User not Login`)
		a.WriteMsg(&tmpRes,EC_USER_NOT_LOGIN.Code())
		return
	}

	if m.Betting < 0 || m.LineCnt < 0 || m.LineCnt > 9 || m.Betting > ExpTable[tmpData.Level].MaxBetting{
		log.Debug(`User Betting Wrong`)
		a.WriteMsg(&tmpRes,EC_WRONG_BETTING.Code())
		return
	}
	if m.GameLevel < 0 || m.GameLevel >= MAX_GAME_LEVEL{
		m.GameLevel = 0
	}
	//扣金币
	NeedGold := m.LineCnt * m.Betting
	err := CostGold(a,NeedGold)
	if err != nil{
		log.Debug(`UpdatePlayCache DB error:[%s]`,err.Error())
		if err.Error() == `Not Enouth Gold`{
			a.WriteMsg(&tmpRes,EC_USER_GOLE_NOT_ENOUGH.Code())
		}else if err.Error() == `Not Login`{
			a.WriteMsg(&tmpRes,EC_USER_NOT_LOGIN.Code())
		}else{
			a.WriteMsg(&tmpRes,EC_DB.Code())
		}
		return
	}
	//增加奖池
	AddCurrentGoldPool(NeedGold,m.GameLevel)
	//檢測是否specialRound
	IsAllZero,IsAll,IsFG,Times,err  := GetGameRoundParam(a,m.GameLevel)
	log.Debug(`User:[%s] StartRound: IsAllZeor:[%v]  IsAll:[%v]  IsFG:[%v]  Times:[%d]`,tmpData.Name,IsAllZero,IsAll,IsFG,Times)
	if IsAllZero{
		tmpRes,err = GeneralAllZero(m.GameLevel,a,m.Betting,m.LineCnt,Times)
		if err != nil{
			log.Debug(`GeneralAllZero error:[%s]`,err.Error())
			a.WriteMsg(&tmpRes,EC_GAME_GENERAL_ERROR.Code())
			return
		}
	}else if IsAll{
		tmpRes,err = GeneralAllButNotZero(m.GameLevel,a,m.Betting,m.LineCnt,Times)
		if err != nil{
			log.Debug(`GeneralAllButNotZero error:[%s]`,err.Error())
			a.WriteMsg(&tmpRes,EC_GAME_GENERAL_ERROR.Code())
			return
		}
	}else if IsFG{
		tmpRes,err = GeneralFC(m.GameLevel,a,m.Betting,m.LineCnt,Times)
		if err != nil{
			log.Debug(`GeneralFC error:[%s]`,err.Error())
			a.WriteMsg(&tmpRes,EC_GAME_GENERAL_ERROR.Code())
			return
		}
	}else {
		tmpRes,err = GeneralNotFC(m.GameLevel,a,m.Betting,m.LineCnt,Times)
		if err != nil{
			log.Debug(`GeneralNotFC error:[%s]`,err.Error())
			a.WriteMsg(&tmpRes,EC_GAME_GENERAL_ERROR.Code())
			return
		}
	}

	err = AddGold(a,tmpRes.TotalScore)
	if err != nil{
		log.Debug(`AddGold  error:[%s]`,err.Error())
	}

	//拿最新的玩家数据进行更新
	tmpData = a.UserData().(User)
	UpdatePlayCache(tmpData)
	a.WriteMsg(&tmpRes,EC_NONE.Code())
	go UserInfoNotice(tmpData.ID,a)
	return
}

