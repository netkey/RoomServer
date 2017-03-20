package internal
import (
	"math/rand"
	"time"
)


type CalTable struct {
	ItemList []int
}


type UserLevelSettings struct {
	UpLevelExp 				int32
	MaxBetting 				int32
}

type UserVipLevelSettings struct {

}


var AllItems []SingleItem
var RuleLines [][]int32

var MarryGameThreeTimes int32
var MarryGameFourTimes int32
var ExpTable []UserLevelSettings
var randomSeed *rand.Rand


func init(){
	LoadRuleLines()
	LoadAllItems()
	LoadExpTable()

}



func LoadRuleLines(){
	RuleLines = append(RuleLines,[]int32{5,6,7,8,9})
	RuleLines = append(RuleLines,[]int32{0,1,2,3,4})
	RuleLines = append(RuleLines,[]int32{10,11,12,13,14})
	RuleLines = append(RuleLines,[]int32{0,6,12,8,4})
	RuleLines = append(RuleLines,[]int32{10,6,2,8,14})
	RuleLines = append(RuleLines,[]int32{0,1,7,3,4})
	RuleLines = append(RuleLines,[]int32{10,11,7,13,14})
	RuleLines = append(RuleLines,[]int32{5,11,12,13,9})
	RuleLines = append(RuleLines,[]int32{5,1,2,3,9})
	//MarryGameTable = []int{10,3,6,8,1,7,6,4,10,5,8,7,2,10,3,6,8,0,7,6,4,10,5,8,7,2}
	//MarryGameTimesTable = []int32{int32(0),int32(200),int32(100),int32(70),int32(50),int32(20),int32(10),int32(5),int32(2)}
	MarryGameThreeTimes = int32(20)
	MarryGameFourTimes = int32(500)
	randomSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func LoadAllItems(){
	AllItems = append(AllItems,SingleItem{
		Id:0,
		AllParam:5000,
		FiveParam:2000,
		FourParam:0,
		ThreeParam:0,
		SpecialParam:0,
		MarryGameBonus:200,
		AppearRate:1,
	})
	AllItems = append(AllItems,SingleItem{
		Id:1,
		AllParam:2500,
		FiveParam:1000,
		FourParam:200,
		ThreeParam:50,
		SpecialParam:0,
		MarryGameBonus:100,
		AppearRate:1,
	})
	AllItems = append(AllItems,SingleItem{
		Id:2,
		AllParam:1000,
		FiveParam:400,
		FourParam:80,
		ThreeParam:20,
		SpecialParam:0,
		MarryGameBonus:40,
		AppearRate:2,
	})
	AllItems = append(AllItems,SingleItem{
		Id:3,
		AllParam:500,
		FiveParam:200,
		FourParam:40,
		ThreeParam:15,
		SpecialParam:50,
		MarryGameBonus:20,
		AppearRate:2,
	})
	AllItems = append(AllItems,SingleItem{
		Id:4,
		AllParam:400,
		FiveParam:160,
		FourParam:30,
		ThreeParam:10,
		SpecialParam:50,
		MarryGameBonus:10,
		AppearRate:3,
	})
	AllItems = append(AllItems,SingleItem{
		Id:5,
		AllParam:250,
		FiveParam:100,
		FourParam:20,
		ThreeParam:7,
		SpecialParam:50,
		MarryGameBonus:8,
		AppearRate:3,
	})
	AllItems = append(AllItems,SingleItem{
		Id:6,
		AllParam:150,
		FiveParam:60,
		FourParam:15,
		ThreeParam:5,
		SpecialParam:15,
		MarryGameBonus:6,
		AppearRate:4,
	})
	AllItems = append(AllItems,SingleItem{
		Id:7,
		AllParam:100,
		FiveParam:40,
		FourParam:10,
		ThreeParam:3,
		SpecialParam:15,
		MarryGameBonus:4,
		AppearRate:4,
	})
	AllItems = append(AllItems,SingleItem{
		Id:8,
		AllParam:50,
		FiveParam:20,
		FourParam:5,
		ThreeParam:2,
		SpecialParam:15,
		MarryGameBonus:2,
		AppearRate:6,
	})
}

func LoadExpTable(){
	for i:=int32(0);i < 101; i++{
		ExpTable = append(ExpTable,UserLevelSettings{
			UpLevelExp:i*int32(1000),
			MaxBetting:i*int32(10),
		})
	}
	return
}