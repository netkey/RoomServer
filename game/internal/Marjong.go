package internal
import (
	"math/rand"
	"time"
	"github.com/name5566/leaf/gate"
	"agent/vendor/github.com/pkg/errors"
	"RoomServer/db"
)

var randomSeed *rand.Rand

func init(){
	randomSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type PlayerCurrentData struct {
	Players 	*gate.Agent
	MJInHand   	[]MJPai
	MJShow	 	[]MJPai
}

type TableGameData struct {
	PlayerInfo       	[]PlayerCurrentData 	//东南西北
	CurrentGameTable 	[]MJPai
	RestInWall		 	[]MJPai
	PutInTable			[]MJPai
	StartIndex		 	int32
	CurrentRound	 	int32
	CurrentPlayerIndex	int32					//东南西北
}


func GeneralRandomMJTable()(returnList []MJPai){
	returnList = make([]int32,0,TotalMarjong)
	for i:= 0 ; i < 4; i++{
		//MJPAI_ZFB
		for j:= 1; j <= 3; j++{
			tmp := MJPai{
				M_Type:MJPAI_ZFB,
				M_Value:j,
			}
			returnList = append(returnList,tmp)
		}
		//MJPAI_FENG
		for j:= 1; j <= 4; j++{
			tmp := MJPai{
				M_Type:MJPAI_FENG,
				M_Value:j,
			}
			returnList = append(returnList,tmp)
		}
		//MJPAI_WAN
		for j:= 1; j <= 9; j++{
			tmp := MJPai{
				M_Type:MJPAI_WAN,
				M_Value:j,
			}
			returnList = append(returnList,tmp)
		}
		//MJPAI_TIAO
		for j:= 1; j <= 9; j++{
			tmp := MJPai{
				M_Type:MJPAI_TIAO,
				M_Value:j,
			}
			returnList = append(returnList,tmp)
		}
		//MJPAI_BING
		for j:= 1; j <= 9; j++{
			tmp := MJPai{
				M_Type:MJPAI_BING,
				M_Value:j,
			}
			returnList = append(returnList,tmp)
		}
	}
	currentIndex := 0
	for i:= 0 ; i < TotalMarjong; i++{
		rand := ((randomSeed.Int31()%(TotalMarjong-currentIndex)) + currentIndex)
		tmp := returnList[currentIndex]
		returnList[currentIndex] = returnList[rand]
		returnList[rand] = tmp
		currentIndex++
	}
	return
}


func DelMJPai(index int,list []MJPai)(res []MJPai,err error){
	if len(list) <= index{
		err = errors.New(`Not Enough MJPai`)
		return
	}
	if index == 0{
		res = list[1:]
	}else if index == len(list) - 1{
		res = list[:index]
	}else {
		res = append(list[0:index],list[index+1:]...)
	}
	return
}




func InitGame(players []*gate.Agent)(resData TableGameData,err error){
	resData = TableGameData{}
	if len(players) != 4{
		err = errors.New(`Not Enough Player`)
		return
	}
	initTable := GeneralRandomMJTable()
	resData.CurrentGameTable = initTable
	resData.RestInWall = append(resData.RestInWall,initTable...)
	//先发牌
	for k,_ := range players{
		tmp := PlayerCurrentData{
			Players:players[k],
		}
		for i := 0; i<MarjonInHand; i++{
			tmp.MJInHand = append(tmp.MJInHand,resData.RestInWall[0])
			resData.RestInWall,_ = DelMJPai(0,resData.RestInWall)
		}
	}
	resData.StartIndex = 0
	resData.CurrentPlayerIndex = 0
	return
}


func PlayerGetMJPai(players *gate.Agent,resData *TableGameData)(err error){
	a := *players
	tmpInData := a.UserData().(db.User)
	b := *resData.PlayerInfo[resData.CurrentPlayerIndex].Players
	tmpTabData := b.UserData().(db.User)
	if tmpInData.ID != tmpTabData.ID{
		err = errors.New(`Wrong player`)
		return
	}
	if len(resData.PlayerInfo[resData.CurrentPlayerIndex].MJInHand) + len(resData.PlayerInfo[resData.CurrentPlayerIndex].MJShow) > MarjonInHand{
		err = errors.New(`Wrong Get`)
		return
	}
	resData.PlayerInfo[resData.CurrentPlayerIndex].MJInHand = append(resData.PlayerInfo[resData.CurrentPlayerIndex].MJInHand,resData.RestInWall[0])
	resData.RestInWall,err = DelMJPai(0,resData.RestInWall)
	//check zimo hu

	return
}


func PlayerPutMJPai(players *gate.Agent,putIndex int32,resData *TableGameData)(err error){
	a := *players
	tmpInData := a.UserData().(db.User)
	b := *resData.PlayerInfo[resData.CurrentPlayerIndex].Players
	tmpTabData := b.UserData().(db.User)
	if tmpInData.ID != tmpTabData.ID{
		err = errors.New(`Wrong player`)
		return
	}
	if len(resData.PlayerInfo[resData.CurrentPlayerIndex].MJInHand) + len(resData.PlayerInfo[resData.CurrentPlayerIndex].MJShow) > MarjonInHand{
		err = errors.New(`Wrong Get`)
		return
	}
	resData.PutInTable = append(resData.PutInTable,resData.PlayerInfo[resData.CurrentPlayerIndex].MJInHand[putIndex])
	resData.PlayerInfo[resData.CurrentPlayerIndex].MJInHand,err = DelMJPai(putIndex,resData.PlayerInfo[resData.CurrentPlayerIndex].MJInHand)
	//check hu

	//check gang

	//check peng

	//check chi

	//round change
	return
}


func CheckHu()(){
	return
}

func CheckGang()(){
	return
}

func CheckPeng()(){
	return
}

func CheckChi()(){
	return
}