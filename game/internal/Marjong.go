package internal
import (
	"math/rand"
	"time"
	"github.com/name5566/leaf/gate"
)

var randomSeed *rand.Rand

func init(){
	randomSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type PlayerCurrentData struct {
	Players 	*gate.Agent
	MJInHand   	[]int32
	MJOnTable 	[]int32
}

type TableGameData struct {
	PlayerInfo       	[]PlayerCurrentData
	CurrentGameTable 	[]int32
	RestList		 	[]int32
	StartIndex		 	int32
	CurrentRound	 	int32
	CurrentPlayerIndex	int32
}


func GeneralRandomMJTable()(returnList []int32){
	returnList = make([]int32,0,TotalMarjong)
	for i:= 0 ; i < 4; i++{
		for j:= 1; j <= TotalType; j++{
			returnList = append(returnList,j)
		}
	}
	currentIndex := 0
	for i:= 0 ; i < 4; i++{
		for j:= 1; j <= TotalType; j++{
			rand := ((randomSeed.Int31()/(TotalMarjong-currentIndex)) + currentIndex)
			tmp := returnList[currentIndex]
			returnList[currentIndex] = returnList[rand]
			returnList[rand] = tmp
		}
	}
	return
}


func InitToEachPlayer()(){
	return
}