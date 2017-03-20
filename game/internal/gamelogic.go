package internal

import (
	"github.com/name5566/leaf/log"
	"RoomServer/msg"
	"github.com/name5566/leaf/gate"
	"errors"
)

const  MAXREPEATTIMES = 1000
const  MAXFCREPEATTIMES = 5000

//玩法相关处理,拉杆,bonus

type SingleItem struct {
	Id		int
	ItemPic		string
	AllParam	int32
	FiveParam	int32
	FourParam	int32
	ThreeParam	int32
	SpecialParam	int32
	MarryGameBonus  int32
	AppearRate	int32 			//越大出现概率越高, single/sum(all single)
}

func GenerateTable()(calTable []int32){
	AllItemSum := int32(0)
	for _,v := range AllItems{
		AllItemSum = AllItemSum + v.AppearRate
	}
	for k := 0; k < 15; k++{
		itemId := GetSingleItem(AllItemSum)
		calTable = append(calTable,int32(itemId))
	}
	return
}

func GetSingleItem(AllItemSum int32)(itemID int32){
	itemID = int32(8)
	randNum := randomSeed.Int31()%AllItemSum
	currentSum := int32(0)
	for _,v := range AllItems{
		currentSum = currentSum + v.AppearRate
		if currentSum >= randNum{
			itemID = int32(v.Id)
			break
		}
	}
	return
}

//计算盘面倍率

type ResultPatten struct {
	RuleLineId	int
	HitCount 	int
	HitItemId  	int
	HitTimes	int32
	SpecialRounds   int
	PattenStr	string //use (XX000)
	SpecialType int
}

func Calculation(calTable []int32,LineCnt int32)(resultMap []msg.PattenData){
	log.Debug("In Calculation: \n[%v]\n[%v]\n[%v]",calTable[0:5],calTable[5:10],calTable[10:15])
	resultMap = make([]msg.PattenData,0)
	//优先检查是否有全盘
	TableList := []int{0,0,0,0,0,0,0,0,0}
	for _,val := range calTable{
		TableList[val]++
	}
	//检查是否15个全部一致
	for index,val2 := range TableList{
		if val2 == 15{
			tmp := msg.PattenData{
				RuleLineId:0,
				HitCount:15,
				HitItemId:int32(index),
				HitTimes:AllItems[index].AllParam,
				SpecialType:1,
				PattenStr:[]int32{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14},
			}
			if index == 0{
				tmp.SpecialRounds = 27
			}
			resultMap = append(resultMap,tmp)
			return
		}
	}

	//检查是否15个存在同类型规则,检查完仍需要按线校验倍率
	if (TableList[3] +TableList[4] + TableList[5]) == 15{
		tmp := msg.PattenData{
			RuleLineId:0,
			HitCount:15,
			HitItemId:3,
			HitTimes:AllItems[3].SpecialParam,
			SpecialType:2,
			PattenStr:[]int32{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14},
		}
		resultMap = append(resultMap,tmp)
	}else if (TableList[6] +TableList[7] + TableList[8]) == 15{
		tmp := msg.PattenData{
			RuleLineId:0,
			HitCount:15,
			HitItemId:6,
			HitTimes:AllItems[6].SpecialParam,
			SpecialType:3,
			PattenStr:[]int32{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14},
		}
		resultMap = append(resultMap,tmp)
	}

	//检查连续相同的元素，id 0 为特殊元素，可以代替所有
	//k为线ID;v为数组中下标,如下图
	//0,  1,  2,  3,  4
	//5,  6,  7,  8,  9
	//10, 11, 12, 13, 14
	for k,v := range RuleLines[0:LineCnt]{
		CheckList:= []int{int(calTable[v[0]]),int(calTable[v[1]]),int(calTable[v[2]]),int(calTable[v[3]]),int(calTable[v[4]])}
		//先找0的个数
		ItemCntList := []int{0,0,0,0,0,0,0,0,0}
		for _,v2 := range CheckList{
			ItemCntList[v2]++
		}
		MaxItemList := []int{
			ItemCntList[0],
			ItemCntList[0]+ItemCntList[1],
			ItemCntList[0]+ItemCntList[2],
			ItemCntList[0]+ItemCntList[3],
			ItemCntList[0]+ItemCntList[4],
			ItemCntList[0]+ItemCntList[5],
			ItemCntList[0]+ItemCntList[6],
			ItemCntList[0]+ItemCntList[7],
			ItemCntList[0]+ItemCntList[8],
		}
		tmpRate := int32(0)
		hitItem := -1
		hitCnt := 0
		var patten  []int32
		specialRound := 0
		//K3代表块ID，V3代表块出现的次数{包含0}
		for k3,v3 := range MaxItemList{
			switch v3{
				case 5:{
					tmpRate = AllItems[k3].FiveParam
					hitItem = k3
					hitCnt = 5
					patten = []int32{v[0],v[1],v[2],v[3],v[4]}
					}
				case 4:{
					//只有一个不同(10000,00001)
					if (CheckList[0] != k3 && CheckList[0] != 0){
						tmpRate = tmpRate + AllItems[k3].FourParam
						hitItem = k3
						hitCnt = 4
						patten = []int32{v[1],v[2],v[3],v[4]}

					}else if (CheckList[4] != k3 && CheckList[4] != 0){
						tmpRate = tmpRate + AllItems[k3].FourParam
						hitItem = k3
						hitCnt = 4
						patten = []int32{v[0],v[1],v[2],v[3]}
					}else if (CheckList[1] != k3 && CheckList[1] != 0) {
						//检测三个连续
						tmpRate = tmpRate + AllItems[k3].ThreeParam
						hitItem = k3
						hitCnt = 3
						patten = []int32{v[2],v[3],v[4]}
					}else if (CheckList[3] != k3 && CheckList[3] != 0){
						//检测三个连续
						tmpRate = tmpRate + AllItems[k3].ThreeParam
						hitItem = k3
						hitCnt = 3
						patten = []int32{v[0],v[1],v[2]}
					}
				}
				case 3:{
					//两个不同 (X000X,XX000,000XX)
					if ( CheckList[1] != k3 && CheckList[1] != 0 && CheckList[0] != k3 && CheckList[0] != 0){
						tmpRate = tmpRate +  AllItems[k3].ThreeParam
						hitItem = k3
						hitCnt = 3
						patten = []int32{v[2],v[3],v[4]}
					}else if (CheckList[3] != k3 && CheckList[3] != 0 && CheckList[4] != k3 && CheckList[4] != 0){
						tmpRate = tmpRate +  AllItems[k3].ThreeParam
						hitItem = k3
						hitCnt = 3
						patten = []int32{v[0],v[1],v[2]}
					}
				}
			}
			//连续0个数确定特别游戏的轮数
			if k3 == 0{
				if v3 == 5{
					specialRound = 3
				}else if v3 == 4 {
					if CheckList[0] != 0 || CheckList[4] != 0{
						specialRound = 2
					}else if CheckList[2] == 0{
						specialRound = 1
					}
				}else if v3 == 3{
					if (CheckList[0] != 0 && CheckList[1] != 0) ||
						(CheckList[3] != 0 && CheckList[4] != 0) {
						specialRound = 1
					}
				}
			}
		}
		if tmpRate > 0{
			tmp := msg.PattenData{
				RuleLineId:int32(k),
				HitCount:int32(hitCnt),
				HitItemId:int32(hitItem),
				HitTimes:tmpRate,
				PattenStr:patten,
				SpecialRounds:int32(specialRound),
			}
			resultMap = append(resultMap,tmp)
		}
	}
	log.Debug(`In Calculation result: [%v]`,resultMap)
	return
}


func GeneralRandomCenter()(cenTable []int32,multParam int32){
	for k := 0; k < 4; k++{
		randNum := randomSeed.Int31()%9
		cenTable = append(cenTable,int32(randNum))
	}
	multParam = 1
	if cenTable[0] == cenTable[1] && cenTable[0] == cenTable[2] && cenTable[0] == cenTable[3]{
		multParam = MarryGameFourTimes
	}else if cenTable[0] == cenTable[1] && cenTable[0] == cenTable[2] {
		multParam = MarryGameThreeTimes
	}else if cenTable[1] == cenTable[2] && cenTable[1] == cenTable[3]{
		multParam = MarryGameThreeTimes
	}
	return
}


func GetMarryGameResult(gameRounds int32,betting int32,fixTimes int32)(result *msg.MarryGameInfos,marryGameTimes int32){
	result = &msg.MarryGameInfos{}
	RoundList := make([]msg.OneRound,gameRounds)
	turnList := make([]msg.SingleTurn,0)
	currentLeft := fixTimes
	for k := 0; k < 10000; k++{
		roundTimes := int32(0)
		tmpTurn := msg.SingleTurn{}
		tmpTable,mul := GeneralRandomCenter()
		roundTimes = roundTimes + mul
		Id := int32(randomSeed.Int31()) % int32(len(AllItems))
		if Id == 0{
			continue
		}else{
			for _,v := range tmpTable{
				if v == Id{
					if AllItems[Id].MarryGameBonus < currentLeft{
						roundTimes = roundTimes + AllItems[Id].MarryGameBonus
						break
					}
				}
			}
			if roundTimes <= currentLeft{
				currentLeft = currentLeft - roundTimes
				tmpTurn.CenterItems = &msg.CenterItems{
					Items:tmpTable,
					MultParam:mul,
				}
				log.Debug(`HitItem:[%d]   !!!!!!!!!! Len:[%d]`,Id,len(AllItems))
				tmpTurn.HitItems = &msg.HitItems{
					HitItems:Id,
					HitBonus:AllItems[Id].MarryGameBonus,
				}
				turnList = append(turnList,tmpTurn)
			}
			if currentLeft < 2{
				break
			}
		}
	}
	marryGameTimes = fixTimes - currentLeft
	for k := 0; k < len(turnList); k++{
		inRoundId := int32(randomSeed.Int31()) % gameRounds
		if turnList[k].HitItems.HitItems == int32(0){
			RoundList[inRoundId].Turns = append(RoundList[inRoundId].Turns,&turnList[k])
			RoundList[inRoundId].Turns = append(RoundList[inRoundId].Turns,&turnList[k+1])
			RoundList[inRoundId].Turns = append(RoundList[inRoundId].Turns,&turnList[k+2])
			RoundList[inRoundId].Turns = append(RoundList[inRoundId].Turns,&turnList[k+3])
			RoundList[inRoundId].Turns = append(RoundList[inRoundId].Turns,&turnList[k+4])
			k = k + 4
		}else{
			RoundList[inRoundId].Turns = append(RoundList[inRoundId].Turns,&turnList[k])
		}
	}
	for _,v := range RoundList{
		result.ResultRounds = append(result.ResultRounds,&v)
	}
	return
}


//產生特定盤面+小遊戲結果

func GeneralAllZero(gameLevel int32,a gate.Agent,betting int32,lineCnt int32,times int32)(resultRes msg.OneArmPlayRes,err error){
	log.Debug(`GeneralAllZero user: [%v]`,a)
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		err = errors.New(`Not Login`)
		return
	}
	//tmpData := a.UserData().(User)
	RepeatTimes := 0
REPEAT_ALL_ZERO:
	resultRes = msg.OneArmPlayRes{
		ShowTable:&msg.PlayTableData{},
		ResultDat:&msg.ResultData{},
		MarryGameResult:&msg.MarryGameInfos{},
	}
	resultRes.ShowTable.TableList = []int32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	tmp := msg.PattenData{
		RuleLineId:0,
		HitCount:15,
		HitItemId:0,
		HitTimes:AllItems[0].AllParam,
		SpecialType:1,
		PattenStr:[]int32{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14},
		SpecialRounds:27,
	}

	resultRes.ResultDat.ResultList = append(resultRes.ResultDat.ResultList,&tmp)
	mainGameTimes := int32(0)
	for _,v :=range resultRes.ResultDat.ResultList{
		mainGameTimes = mainGameTimes + v.HitTimes
	}

	resultRes.MarryGameRounds  = tmp.SpecialRounds
	resultRes.MainGameScore = betting * mainGameTimes

	//Add MarryGame Data  times fix to ?? (to check)
	resultRes.MarryGameResult,resultRes.MarryGameTimes = GetMarryGameResult(resultRes.MarryGameRounds,betting,500)
	resultRes.TotalScore = resultRes.MainGameScore * resultRes.MarryGameTimes
	tmpChk := RTPRateCheck{
		A:a,
		TotalWin:resultRes.TotalScore,
		GameLevel:gameLevel,
		TotalBetting:betting*lineCnt,
	}
	IsOk := CheckRTPRate(tmpChk)
	if !IsOk{
		RepeatTimes = RepeatTimes +1
		if RepeatTimes < MAXREPEATTIMES{
			goto REPEAT_ALL_ZERO
		}else{
			return GeneralNotFC(gameLevel,a,betting,lineCnt,times)
		}
	}
	log.Debug(`GeneralAllZero result: [%v]`,resultRes)
	return
}

func GeneralAllButNotZero(gameLevel int32,a gate.Agent,betting int32,lineCnt int32,times int32)(resultRes msg.OneArmPlayRes,err error){
	log.Debug(`GeneralAllButNotZero user: [%v]`,a)
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		err = errors.New(`Not Login`)
		return
	}
	//tmpData := a.UserData().(User)
	RepeatTimes := 0
REPEAT_ALL_BUT_NOT_ZERO:
	resultRes = msg.OneArmPlayRes{
		ShowTable:&msg.PlayTableData{},
		ResultDat:&msg.ResultData{},
		MarryGameResult:&msg.MarryGameInfos{},
	}
	index := int32(0)
	for k,v :=range AllItems{
		if  v.AllParam - times < 0 {
			index = int32(k)
			break
		}
	}

	resultRes.ShowTable.TableList = []int32{index,index,index,index,index,index,index,index,index,index,index,index,index,index,index}
	tmp :=  msg.PattenData{
		RuleLineId:0,
		HitCount:15,
		HitItemId:index,
		HitTimes:AllItems[index].AllParam,
		SpecialType:1,
		PattenStr:[]int32{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14},
		SpecialRounds:0,
	}

	resultRes.ResultDat.ResultList = append(resultRes.ResultDat.ResultList,&tmp)
	mainGameTimes := int32(0)
	for _,v :=range resultRes.ResultDat.ResultList{
		mainGameTimes = mainGameTimes + v.HitTimes
	}

	resultRes.MarryGameRounds  = tmp.SpecialRounds
	resultRes.MainGameScore = betting * mainGameTimes
	resultRes.TotalScore = resultRes.MainGameScore
	tmpChk := RTPRateCheck{
		A:a,
		TotalWin:resultRes.TotalScore,
		GameLevel:gameLevel,
		TotalBetting:betting*lineCnt,
	}
	IsOk := CheckRTPRate(tmpChk)
	if !IsOk{
		RepeatTimes = RepeatTimes +1
		if RepeatTimes < MAXREPEATTIMES{
			goto REPEAT_ALL_BUT_NOT_ZERO
		}else{
			return GeneralNotFC(gameLevel,a,betting,lineCnt,times)
		}
	}
	log.Debug(`GeneralAllButNotZero result: [%v]`,resultRes)
	return
}


func GeneralNotFC(gameLevel int32,a gate.Agent,betting int32,lineCnt int32,times int32)(resultRes msg.OneArmPlayRes,err error){
	log.Debug(`GeneralNotFC user: [%v]`,a)
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		err = errors.New(`Not Login`)
		return
	}
	//tmpData := a.UserData().(User)
	RepeatTimes := 0
REPEAT_NOT_FC:
	resultRes = msg.OneArmPlayRes{
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
	resultRes.TotalScore = resultRes.MainGameScore
	if SpecialRounds > 0{
		RepeatTimes = RepeatTimes +1
		goto REPEAT_NOT_FC
	}

	tmpChk := RTPRateCheck{
		A:a,
		TotalWin:resultRes.TotalScore,
		GameLevel:gameLevel,
		TotalBetting:betting*lineCnt,
	}
	IsOk := CheckRTPRate(tmpChk)
	if !IsOk{
		RepeatTimes = RepeatTimes +1
		if RepeatTimes < MAXREPEATTIMES{
			goto REPEAT_NOT_FC
		}else{
			return ReturnFixResult()
		}
	}
	log.Debug(`GeneralNotFC result: [%v]`,resultRes)
	return
}

func GeneralFC(gameLevel int32,a gate.Agent,betting int32,lineCnt int32,times int32)(resultRes msg.OneArmPlayRes,err error){
	log.Debug(`GeneralFC user: [%v]`,a)
	if a.UserData() == nil{
		log.Debug(`User not Login`)
		err = errors.New(`Not Login`)
		return
	}
	//tmpData := a.UserData().(User)
	RepeatTimes := 0
	REPEAT_FC:
	resultRes = msg.OneArmPlayRes{
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
		if RepeatTimes < MAXFCREPEATTIMES * 10{
			goto REPEAT_FC
		}else{
			log.Debug(`MAXREPEATTIMES ~~~~ times: [%d]`,times)
			return GeneralNotFC(gameLevel,a,betting,lineCnt,times)
		}
	}
	log.Debug(`MarryGameTimes  times: [%d]`,RepeatTimes)
	//Add MarryGame Data  times fix to ?? (to check)
	resultRes.MarryGameResult,resultRes.MarryGameTimes = GetMarryGameResult(resultRes.MarryGameRounds,betting,times)
	resultRes.TotalScore = resultRes.MainGameScore + betting*lineCnt*resultRes.MarryGameTimes
	tmpChk := RTPRateCheck{
		A:a,
		TotalWin:resultRes.TotalScore,
		GameLevel:gameLevel,
		TotalBetting:betting*lineCnt,
	}
	IsOk := CheckRTPRate(tmpChk)
	if !IsOk{
		RepeatTimes = RepeatTimes + 1
		if RepeatTimes < MAXFCREPEATTIMES * 10{
			goto REPEAT_FC
		}else{
			return GeneralNotFC(gameLevel,a,betting,lineCnt,times)
		}
	}
	log.Debug(`GeneralFC result: [%v]`,resultRes)
	return
}


func ReturnFixResult()(resultRes msg.OneArmPlayRes,err error){
	//return a fix no hit table~
	resultRes = msg.OneArmPlayRes{
		ShowTable:&msg.PlayTableData{},
		ResultDat:&msg.ResultData{},
		MarryGameResult:&msg.MarryGameInfos{},
	}
	resultRes.ShowTable.TableList = []int32{
		5, 2, 7, 4, 0,
		2, 1, 1, 1, 6,
		1, 5, 2, 7, 6,
	}
	return
}

