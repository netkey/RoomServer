package internal
import (
	"RoomServer/msg"
	"github.com/name5566/leaf/gate"
	"strings"
	"github.com/name5566/leaf/log"

	"RoomServer/common"
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
				a.WriteMsg(&tmpRes,common.EC_PARAM.Code())
				return
			}
		}

	}
	a.WriteMsg(&tmpRes,common.EC_NONE.Code())
	return
}
