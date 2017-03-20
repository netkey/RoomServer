package internal
import (
	"RoomServer/msg"
	"github.com/name5566/leaf/gate"
	"strings"
	"github.com/name5566/leaf/log"
	"strconv"
)


func handleGmCmd(args []interface{})(){
	m := args[0].(*msg.GMGameReq)
	a := args[1].(gate.Agent)
	tmpRes := msg.GMGameRes{}
	cmd := m.GmCmd
	cmdParams := strings.Split(cmd,` `)
	switch cmdParams[0] {
		case `GetTable`:{
			if len(cmdParams) < 2{
				log.Debug("Wrong Param: [%s]", cmd)
				a.WriteMsg(&tmpRes,EC_PARAM.Code())
				return
			}
			if cmdParams[1] == `0`{
				go GeneralGMFC(0,a,1,9)
			}else if  cmdParams[1] == `1`{
				if len(cmdParams) >= 3{
					itemId,_ := strconv.Atoi(cmdParams[3])
					go GeneralGMAllNotZero(0,a,1,9,int32(itemId))
				}else{
					go GeneralGMAllNotZero(0,a,1,9,0)
				}

			}else if cmdParams[1] == `2`{
				if len(cmdParams) >= 3{
					itemId,_ := strconv.Atoi(cmdParams[3])
					go GeneralGMSpecialType(0,a,1,9,int32(itemId))
				}else{
					go GeneralGMSpecialType(0,a,1,9,3)
				}
			}
		}

	}
	a.WriteMsg(&tmpRes,EC_NONE.Code())
	return
}



func GeneralGMFC(gameLevel int32,a gate.Agent,betting int32,lineCnt int32)(){
	log.Debug(`GeneralGMFC user: [%v]`,a)
	//tmpData := a.UserData().(User)
	RepeatTimes := 0
	REPEAT_FC:
	resultRes := msg.OneArmPlayRes{
		ShowTable:&msg.PlayTableData{},
		ResultDat:&msg.ResultData{},
		MarryGameResult:&msg.MarryGameInfos{},
	}
	resultRes.ShowTable.TableList = GenerateTable()
	resultData := Calculation(resultRes.ShowTable.TableList,lineCnt)
	TotalTimes:= int32(0)
	SpecialRounds := int32(0)
	for k,v := range resultData{
		resultRes.ResultDat.ResultList = append(resultRes.ResultDat.ResultList,&resultData[k])
		TotalTimes = TotalTimes + int32(v.HitTimes)
		SpecialRounds = SpecialRounds + int32(v.SpecialRounds)
	}

	resultRes.MainGameScore = TotalTimes*betting
	resultRes.MarryGameRounds = SpecialRounds
	if SpecialRounds <= 0{
		RepeatTimes = RepeatTimes +1
		if RepeatTimes < 10000{
			goto REPEAT_FC
		}else{
			log.Debug(`MAXREPEATTIMES ~~~~ times: [%d]`)
			a.WriteMsg(resultRes,EC_NONE.Code())
			return
		}
	}
	log.Debug(`MarryGameTimes  times: [%d]`,RepeatTimes)
	resultRes.MarryGameResult,resultRes.MarryGameTimes = GetMarryGameResult(resultRes.MarryGameRounds,betting,50)
	resultRes.TotalScore = resultRes.MainGameScore + betting*lineCnt*resultRes.MarryGameTimes
	a.WriteMsg(resultRes,EC_NONE.Code())
	log.Debug(`GeneralGMFC result: [%v]`,resultRes)
	return
}

func GeneralGMAllNotZero(gameLevel int32,a gate.Agent,betting int32,lineCnt int32,itemID int32)(){
	log.Debug(`GeneralGMAllNotZero user: [%v]`,a)
	//tmpData := a.UserData().(User)
	resultRes := msg.OneArmPlayRes{
		ShowTable:&msg.PlayTableData{},
		ResultDat:&msg.ResultData{},
		MarryGameResult:&msg.MarryGameInfos{},
	}
	resultRes.ShowTable.TableList = []int32{itemID,itemID,itemID,itemID,itemID,itemID,itemID,itemID,itemID,itemID,itemID,itemID,itemID,itemID,itemID}
	tmp := msg.PattenData{
		RuleLineId:0,
		HitCount:15,
		HitItemId:itemID,
		HitTimes:AllItems[itemID].AllParam,
		SpecialType:1,
		PattenStr:[]int32{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14},
		SpecialRounds:0,
	}
	if itemID == 0{
		tmp.SpecialRounds = 27
	}
	resultRes.ResultDat.ResultList = append(resultRes.ResultDat.ResultList,&tmp)
	mainGameTimes := int32(0)
	for _,v :=range resultRes.ResultDat.ResultList{
		mainGameTimes = mainGameTimes + v.HitTimes
	}
	resultRes.MarryGameRounds  = tmp.SpecialRounds
	resultRes.MainGameScore = betting * mainGameTimes
	//Add MarryGame Data  times fix to ?? (to check)
	resultRes.MarryGameResult,resultRes.MarryGameTimes = GetMarryGameResult(resultRes.MarryGameRounds,betting,50)
	resultRes.TotalScore = resultRes.MainGameScore * resultRes.MarryGameTimes
	a.WriteMsg(resultRes,EC_NONE.Code())
	log.Debug(`GeneralGMAllNotZero result: [%v]`,resultRes)
	return
}

func GeneralGMSpecialType(gameLevel int32,a gate.Agent,betting int32,lineCnt int32,itemID int32)(){
	log.Debug(`GeneralGMSpecialType user: [%v]`,a)

	//tmpData := a.UserData().(User)
	resultRes := msg.OneArmPlayRes{
		ShowTable:&msg.PlayTableData{},
		ResultDat:&msg.ResultData{},
		MarryGameResult:&msg.MarryGameInfos{},
	}
	if itemID != 3 {
		resultRes.ShowTable.TableList = []int32{6,7,8,8,7,6,7,6,7,7,8,7,8,6,6}
	}else {
		resultRes.ShowTable.TableList = []int32{3,3,3,4,3,4,4,3,4,3,4,3,4,3,4}
	}

	tmp := msg.PattenData{
		RuleLineId:0,
		HitCount:15,
		HitItemId:3,
		HitTimes:AllItems[itemID].SpecialParam,
		SpecialType:2,
		PattenStr:[]int32{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14},
		SpecialRounds:0,
	}
	if itemID != 3{
		tmp.SpecialType = 3
	}

	resultRes.ResultDat.ResultList = append(resultRes.ResultDat.ResultList,&tmp)
	mainGameTimes := int32(0)
	for _,v :=range resultRes.ResultDat.ResultList{
		mainGameTimes = mainGameTimes + v.HitTimes
	}
	resultRes.MarryGameRounds  = tmp.SpecialRounds
	resultRes.MainGameScore = betting * mainGameTimes
	//Add MarryGame Data  times fix to ?? (to check)
	resultRes.MarryGameResult,resultRes.MarryGameTimes = GetMarryGameResult(resultRes.MarryGameRounds,betting,50)
	resultRes.TotalScore = resultRes.MainGameScore * resultRes.MarryGameTimes
	a.WriteMsg(resultRes,EC_NONE.Code())
	log.Debug(`GeneralGMSpecialType result: [%v]`,resultRes)
	return
}