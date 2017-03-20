package main

import (
	"RoomServer/msg"
	"encoding/binary"
	"net"
	"io"
	"fmt"
	"github.com/golang/protobuf/proto"
	//"github.com/gobs/pretty"
	"math/rand"
	"time"
	"github.com/gobs/pretty"
)

type  PlayerData struct {
	PlayTimes 		int
	UserData msg.UserData
	TotalBetting	int32
	MarryGameTimes  int32
	MaxWin			int32
}


var currentPlayer PlayerData
var getGold = int32(0)
var randomSeed *rand.Rand
var receiveTime = 0
var RateMap  = make(map[int32]int)


func main() {
	conn, err := net.Dial("tcp", "104.224.130.90:3544")
	//conn, err := net.Dial("tcp", "127.0.0.1:3544")

	randomSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}
	cmd := uint16(9999)
	fmt.Print(`SendCmd:`)
	go AcceptAns(&conn)
	for {
		fmt.Scanf("%d", &cmd)
		if cmd <10000{
			SendingReq(cmd,&conn)
		}
	}

	return
}

const BufLength = 1024

func AcceptAns(conn *net.Conn){
	TcpCon := *conn
	for {
		data := make([]byte, 0)//此处做一个输入缓冲以免数据过长读取到不完整的数据
		buf := make([]byte, BufLength)
		n, err := TcpCon.Read(buf)

		if err != nil && err != io.EOF {
			println(err.Error())
			return
		}
		data = append(data, buf[:n]...)
		Length := binary.BigEndian.Uint16(data[0:2])
		Body := data[2:Length+2]
		cmd := binary.BigEndian.Uint16(Body[0:2])
		errorCode := binary.BigEndian.Uint16(Body[2:4])
		switch cmd {
		case uint16(msg.MsgID_loginRes):
			{
				fmt.Printf("Receive Cmd:[%d]! Code:[%d]\n",cmd,errorCode)
				for _,v :=range data{
					fmt.Printf("[%d]",v)
				}
				fmt.Printf("\n")
				tmp := msg.LoginRes{}
				err := proto.Unmarshal(Body[4:],&tmp)
				if err != nil{
					fmt.Printf(`err:%s`,err.Error())
				}
				fmt.Printf("Receive Package:%v \n",tmp)
			}
		case uint16(msg.MsgID_oneArmPlayRes):
			{
				fmt.Printf("Receive Cmd:[%d]! Code:[%d]\n",cmd,errorCode)
				/*for _,v :=range data{
					fmt.Printf("[%d]",v)
				}
				fmt.Printf("\n")*/
				tmp := msg.OneArmPlayRes{}
				err := proto.Unmarshal(Body[4:],&tmp)
				if err != nil{
					fmt.Printf(`err:%s`,err.Error())
				}
				if tmp.MarryGameRounds > 0{
					currentPlayer.MarryGameTimes = currentPlayer.MarryGameTimes + 1
				}
				getGold = getGold + tmp.TotalScore
				if tmp.TotalScore > currentPlayer.MaxWin{
					currentPlayer.MaxWin = tmp.TotalScore
				}
				pretty.PrettyPrint(tmp)
				receiveTime = receiveTime + 1
				RateMap[tmp.TotalScore] = RateMap[tmp.TotalScore] +1
				fmt.Printf(`pt:[%d]`,receiveTime )
				fmt.Printf("Request:TotalWin:[%d]  MainGameWin:[%d] MarryGame:[%d]  MarryTimes:[%d]\n",tmp.TotalScore,tmp.MainGameScore,tmp.MarryGameRounds,tmp.MarryGameTimes)
			}
		case uint16(msg.MsgID_playerInfoUpdate):
			{
				fmt.Printf("Receive Cmd:[%d]! Code:[%d]\n",cmd,errorCode)
				tmp := msg.PlayerInfoUpdate{}
				err := proto.Unmarshal(Body[4:],&tmp)
				if err != nil{
					fmt.Printf(`err:%s`,err.Error())
				}
				//pretty.PrettyPrint(tmp)
				currentPlayer.UserData = *tmp.UserInfo
				fmt.Printf("Current Gold:[%d]\n",currentPlayer.UserData.Gold)
			}
		/*case uint16(msg.MsgID_marryGameRes):
			{
				fmt.Printf("Receive Cmd:[%d]! Code:[%d]\n",cmd,errorCode)
				tmp := msg.MarryGameRes{}
				err := proto.Unmarshal(Body[4:],&tmp)
				if err != nil{
					fmt.Printf(`err:%s`,err.Error())
				}
				pretty.PrettyPrint(tmp)
			}*/
		}
	}
}


func SendingReq(Cmd uint16,conn *net.Conn){
	switch Cmd {
		case uint16(msg.MsgID_loginReq):{
			UserName := ``
			PassWd := ``
			LogType := int32(1)
			AccessToken := `test`
			fmt.Print(`UserName:`)
			fmt.Scanln(&UserName)
			fmt.Print(`PassWd:`)
			fmt.Scanln(&PassWd)
			tmp := msg.LoginReq{
				UserName:UserName,
				PassWd:PassWd,
				LoginType:LogType,
				AccessToken:AccessToken,
			}
			data, err := proto.Marshal(&tmp)
			if err != nil {
				println(err.Error())
				return
			}
			//fmt.Printf("DataLength:%d \n",len(data))
			totalLen := 4 + len(data)
			m := make([]byte, totalLen)
			binary.BigEndian.PutUint16(m, uint16(len(data) + 2))
			binary.BigEndian.PutUint16(m[2:], (Cmd))
			copy(m[4:], data)
			/*for _, v := range m {
				fmt.Printf("[%d]", v)
			}*/
			fmt.Printf("\n")
			// 发送消息
			ConnTcp := *conn
			ConnTcp.Write(m)
		}
		case uint16(msg.MsgID_oneArmPlayReq):{
			Lines := int32(0)
			Betting := int32(0)
			fmt.Print(`Lines:`)
			fmt.Scanln(&Lines)
			fmt.Print(`Betting:`)
			fmt.Scanln(&Betting)
			tmp := msg.OneArmPlayReq{
				LineCnt:Lines,
				Betting:Betting,
			}
			data, err := proto.Marshal(&tmp)
			if err != nil {
				println(err.Error())
				return
			}
			//fmt.Printf("DataLength:%d \n",len(data))
			totalLen := 4 + len(data)
			m := make([]byte, totalLen)
			binary.BigEndian.PutUint16(m, uint16(len(data) + 2))
			binary.BigEndian.PutUint16(m[2:], Cmd)
			copy(m[4:], data)
			/*for _, v := range m {
				fmt.Printf("[%d]", v)
			}*/
			currentPlayer.TotalBetting = currentPlayer.TotalBetting + Betting * Lines
			currentPlayer.PlayTimes = currentPlayer.PlayTimes + 1

			fmt.Printf("\n")
			// 发送消息
			ConnTcp := *conn
			ConnTcp.Write(m)
		}
	case uint16(9999):{
		repeatTimes := 0
		fmt.Print(`RepeqtTimes:`)
		fmt.Scanln(&repeatTimes)
		for i := 0; i < repeatTimes; i++{
			Betting := randomSeed.Int31() % 10
			Betting = 1
			Lines := int32(9)
			tmp := msg.OneArmPlayReq{
				LineCnt:Lines,
				Betting:Betting,
			}
			if currentPlayer.UserData.Gold < Lines * Betting{
				break
			}

			currentPlayer.PlayTimes = currentPlayer.PlayTimes +1
			data, err := proto.Marshal(&tmp)
			if err != nil {
				println(err.Error())
				return
			}

			totalLen := 4 + len(data)
			m := make([]byte, totalLen)
			binary.BigEndian.PutUint16(m, uint16(len(data) + 2))
			binary.BigEndian.PutUint16(m[2:], uint16(msg.MsgID_oneArmPlayReq))
			copy(m[4:], data)

			currentPlayer.TotalBetting = currentPlayer.TotalBetting + Betting * Lines
			currentPlayer.PlayTimes = currentPlayer.PlayTimes + 1
			fmt.Printf("Request:Betting:[%d]  Line:[%d]\n",Betting,Lines)
			// 发送消息
			ConnTcp := *conn
			ConnTcp.Write(m)
			time.Sleep(time.Second * 1)
		}

		for  receiveTime < repeatTimes{
			time.Sleep(time.Microsecond * 500)
		}
		fmt.Printf("END: Times:[%d]  UseGold:[%d]  GetGold:[%d] MarryGameTimes:[%d] MaxWin:[%d]\n",repeatTimes,currentPlayer.TotalBetting ,(getGold),currentPlayer.MarryGameTimes,currentPlayer.MaxWin)
		fmt.Printf("ResultRates: [%v]",RateMap)
	}
	}

}

