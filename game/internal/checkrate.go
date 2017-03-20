package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"errors"
	"sync"
)

//概率均以10000为底

//总体平均回报率 X/100
var RTPRate 				[]int32
var CurrentGoldPool			int64
var RTPLock 				*sync.Cond


//当前实现概率方案为 1/平均值


//MarryGame相关概率设置
//小游戏平均频率
var FGAvgRound				[]int32
//小游戏平均倍数
var FGAvgTimes				[]int32


//最少经过多少局出现小游戏
//var FGRoundLowerBound		int32
//最多经过多少局出现小游戏
//var FGRoundUpperBound		int32
//小游戏最低倍数
//var FGTimesLowerBound		int32
//小游戏最高倍数
//var FGTimesUpperBound		int32




//MainGame相关概率设置
//全盘频率(不包含全0)
var AllTableAvgRound 			[]int32
var AllTableAvgTimes			[]int32

//全盘0频率
var AllTableZeroAvgRound 		[]int32


const MAX_GAME_LEVEL = 3

func init(){
	RTPRate = []int32{95,95,95}
	FGAvgRound = []int32{250,250,250}
	FGAvgTimes = []int32{50,50,50}
	AllTableAvgRound = []int32{1000,1300,1500}
	AllTableAvgTimes = []int32{200,200,200}
	AllTableZeroAvgRound = []int32{500000,800000,1000000}
	CurrentGoldPool = 0
	RTPLock = sync.NewCond(new(sync.Mutex))
}




func GetGameRoundParam(a gate.Agent,GameLevel int32)(IsAllZero bool,IsAll bool,IsFG bool,Times int32,err error){
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		err = errors.New(`Not Login`)
		return
	}

	if GameLevel >= MAX_GAME_LEVEL || GameLevel < 0{
		log.Debug(`GameLevel Error,Input:%d`,GameLevel)
		err = errors.New(`GameLevel Error`)
		return
	}
	//tmpData := a.UserData().(User)
	//check user can use this level
	//
	AllZeroRand := randomSeed.Int31()%(AllTableZeroAvgRound[GameLevel])
	if AllZeroRand == 1{
		Times = 50
		IsAllZero = true
		return
	}

	AllTableRand := randomSeed.Int31()%(AllTableAvgRound[GameLevel])
	if AllTableRand == 1{
		Times = 500
		IsAll = true
		return
	}

	FGRand := randomSeed.Int31()%(FGAvgRound[GameLevel])
	if FGRand == 1{
		Times = 50
		IsFG = true
		return
	}
	return
}
/*
func CheckRTPRate(a gate.Agent,totalWin int32,GameLevel int32)(Ok bool,err error){
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		err = errors.New(`Not Login`)
		return
	}
	tmpData := a.UserData().(User)
	//還需要添加首充,充值,新玩家檢測
	Ok = false
	MaxReturn := tmpData.TotalBetting * int64(RTPRate[GameLevel]) / 100  + (1000)   //基礎值1000,保證起步
	if MaxReturn >= int64(totalWin){
		Ok = true
		return
	}
	return
}*/
type RTPRateCheck struct {
	A 			gate.Agent
	TotalBetting		int32
	TotalWin 		int32
	GameLevel 		int32
}



func CheckRTPRate(checkData RTPRateCheck)(bool){
	log.Debug(`Current Gold Pool:%d`,CurrentGoldPool)
	RTPLock.L.Lock()
	log.Debug(`Get Check:%v`,checkData)
	if checkData.A.UserData() == nil{
		log.Debug(`User not Login`)
		RTPLock.L.Unlock()
		return false
	}
	//tmpData := i.A.UserData().(User)

	if int64(checkData.TotalWin) > CurrentGoldPool{
		RTPLock.L.Unlock()
		return false
	}

	CurrentGoldPool =  CurrentGoldPool - int64(checkData.TotalWin)
	RTPLock.L.Unlock()

	return true
}

func AddCurrentGoldPool(AddGold int32,GameLevel int32){
	RTPLock.L.Lock()
	CurrentGoldPool = CurrentGoldPool + (int64(AddGold) *int64(RTPRate[GameLevel]) / int64(100))
	RTPLock.L.Unlock()
}
