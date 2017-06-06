package db
import (
	"github.com/name5566/leaf/log"
	"errors"
	"RoomServer/msg"
	"github.com/name5566/leaf/gate"
	"strconv"
	"RoomServer/conf"
	"time"
	"RoomServer/common"
)

//处理player的数据逻辑，vip系统，等级成长，金币

type User struct {
	ID 					string
	Name 				string
	Gold				int32
	Exp					int32
	Diamond 			int32
	Avatar				string
	Mobile				string
	PassWd				string
	FbToken				string
	VipLevel 			int32
	Level				int32
	LastLoginTime   	time.Time
	IsLevelUp			bool
	IsVipLevelUp		bool
	TotalBetting 		int64
	TotalGame			int64
	PlayerRoomID		string
}



func GetPlayerByUserName(Name string)(userData User,err error){
	log.Debug(`GetPlayerByUserName:[%s]`,Name)
	sqll := `select id,name,gold,exp,diamond,avatar,mobile,passwd,fbtoken,vip_level,player_level from user where name = ?`
	row,err := RAWDB.Query(sqll,Name)
	if err != nil{
		log.Debug(`Db query err:%v`,err.Error())
		return
	}
	if row.Next(){
		tmpUser := User{}
		row.Scan(&tmpUser.ID,&tmpUser.Name,&tmpUser.Gold,&tmpUser.Exp,&tmpUser.Diamond,&tmpUser.Avatar,&tmpUser.Mobile,&tmpUser.PassWd,&tmpUser.FbToken,&tmpUser.VipLevel,&tmpUser.Level)
		userData = tmpUser
	}else {
		if conf.Server.IsStg{
			backId,errCr := CreateUser(Name,Name)
			if errCr == nil{
				userData,err  = GetPlayerByUserId(strconv.FormatInt(backId,10))
				return
			}
		}
		log.Debug(`No This User:[%s]`,Name)
		err = errors.New(`No This User`)
		return
	}
	log.Debug(`GetUser:[%s]`,Name)
	return
}


func GetPlayerByUserId(id string)(userData User,err error){
	sqll := `select id,name,gold,exp,diamond,avatar,mobile,passwd,fbtoken,vip_level,player_level from user where id = ?`
	row,err := RAWDB.Query(sqll,id)
	if err != nil{
		log.Debug(`Db query err:%v`,err.Error())
		return
	}
	if row.Next(){
		tmpUser := User{}
		row.Scan(&tmpUser.ID,&tmpUser.Name,&tmpUser.Gold,&tmpUser.Exp,&tmpUser.Diamond,&tmpUser.Avatar,&tmpUser.Mobile,&tmpUser.PassWd,&tmpUser.FbToken,&tmpUser.VipLevel,&tmpUser.Level)
		userData = tmpUser
	}else{
		err = errors.New(`No This User`)
	}
	return
}


func UserInit(userInfo *User){
	//config

}

func UserInfoNotice(userId string,a gate.Agent)(){
	tmpData := a.UserData().(User)
	IdInt ,_:= strconv.ParseInt(tmpData.ID,10,32)
	tmpRes := msg.PlayerInfoUpdate{
		UserInfo:&msg.UserData{
			UserId:int32(IdInt),
			UserName:tmpData.Name,
			Avatar:tmpData.Avatar,
			Gold:tmpData.Gold,
			Exp:tmpData.Exp,
			Diamond:tmpData.Diamond,
			Level:tmpData.Level,
			VipLevel:tmpData.VipLevel,
		},
		ChangeMsg:&msg.ChangeMsg{
			IsLevelUp:tmpData.IsLevelUp,
			IsVipLevelUp:tmpData.IsVipLevelUp,
			IsHaveNewMail:false,
		},
	}
	tmpData.IsLevelUp = false
	tmpData.IsVipLevelUp = false
	a.SetUserData(tmpData)
	a.WriteMsg(&tmpRes,common.EC_NONE.Code())
	return
}


func CreateUser(name string,passwd string)(backID int64,err error){
	sqll := `insert into user (name,passwd,gold) values (?,?,?)`
	result,err := RAWDB.Exec(sqll,name,passwd,100)
	if err != nil{
		log.Debug(`Db CreateUser err:%v`,err.Error())
		return
	}
	backID,err = result.LastInsertId()
	if err != nil{
		log.Debug(`Db CreateUser err:%v`,err.Error())
		return
	}
	return
}

func UpdatePlayCache(userCache User)(err error){
	sqll := `update user set gold = ?,exp = ?,diamond = ?,avatar = ?,mobile = ?,player_level = ?,vip_level = ? where id = ?`
	result,err := RAWDB.Exec(sqll,userCache.Gold,userCache.Exp,userCache.Diamond,userCache.Avatar,userCache.Mobile,userCache.Level,userCache.VipLevel,userCache.ID)
	if err != nil{
		log.Debug(`Db UpdatePlayCache err:%v`,err.Error())
		return
	}
	affRow,err := result.RowsAffected()
	if err != nil || affRow == 0{
		err = errors.New(`No This User`)
	}
	return
}




func CostGold(a gate.Agent,costGold int32)(err error){
	tmpData := a.UserData().(User)
	if tmpData.ID == ``{
		log.Debug(`User not Login`)
		err = errors.New(`Not Login`)
		return
	}
	if costGold > tmpData.Gold{
		log.Debug(`User Gold Not Enough`)
		err = errors.New(`Not Enouth Gold`)
		return
	}
	tmpData.Gold = tmpData.Gold - costGold
	tmpData.Exp = tmpData.Exp + costGold        						//未处理VIP
	tmpData.TotalBetting = tmpData.TotalBetting + int64(costGold)		//記錄總投注數
	a.SetUserData(tmpData)
	return
}


func AddGold(a gate.Agent,AddGold int32)(err error){
	tmpData := a.UserData().(User)
	if tmpData.ID == ``{
		log.Debug(`User not Login`)
		err = errors.New(`Not Login`)
		return
	}
	tmpData.Gold = tmpData.Gold + AddGold
	a.SetUserData(tmpData)
	return
}


func UpdatePass(id string,pass int)(err error){
	sqll := `update user set passwd = ? where id = ?`
	result,err := RAWDB.Exec(sqll,pass,id)
	if err != nil{
		log.Debug(`Db UpdatePass err:%v`,err.Error())
		return
	}
	affRow,err := result.RowsAffected()
	if err != nil || affRow == 0{
		err = errors.New(`No This User`)
	}
	return
}
