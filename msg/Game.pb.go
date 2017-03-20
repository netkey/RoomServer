// Code generated by protoc-gen-go.
// source: Game.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// /////////////////主玩法////////////////////////
type PlayTableData struct {
	TableList []int32 `protobuf:"varint,1,rep,packed,name=tableList" json:"tableList,omitempty"`
}

func (m *PlayTableData) Reset()                    { *m = PlayTableData{} }
func (m *PlayTableData) String() string            { return proto.CompactTextString(m) }
func (*PlayTableData) ProtoMessage()               {}
func (*PlayTableData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *PlayTableData) GetTableList() []int32 {
	if m != nil {
		return m.TableList
	}
	return nil
}

type PattenData struct {
	RuleLineId    int32   `protobuf:"varint,1,opt,name=ruleLineId" json:"ruleLineId,omitempty"`
	HitCount      int32   `protobuf:"varint,2,opt,name=hitCount" json:"hitCount,omitempty"`
	HitItemId     int32   `protobuf:"varint,3,opt,name=hitItemId" json:"hitItemId,omitempty"`
	HitTimes      int32   `protobuf:"varint,4,opt,name=hitTimes" json:"hitTimes,omitempty"`
	SpecialRounds int32   `protobuf:"varint,5,opt,name=specialRounds" json:"specialRounds,omitempty"`
	PattenStr     []int32 `protobuf:"varint,6,rep,packed,name=pattenStr" json:"pattenStr,omitempty"`
	SpecialType   int32   `protobuf:"varint,7,opt,name=specialType" json:"specialType,omitempty"`
}

func (m *PattenData) Reset()                    { *m = PattenData{} }
func (m *PattenData) String() string            { return proto.CompactTextString(m) }
func (*PattenData) ProtoMessage()               {}
func (*PattenData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PattenData) GetRuleLineId() int32 {
	if m != nil {
		return m.RuleLineId
	}
	return 0
}

func (m *PattenData) GetHitCount() int32 {
	if m != nil {
		return m.HitCount
	}
	return 0
}

func (m *PattenData) GetHitItemId() int32 {
	if m != nil {
		return m.HitItemId
	}
	return 0
}

func (m *PattenData) GetHitTimes() int32 {
	if m != nil {
		return m.HitTimes
	}
	return 0
}

func (m *PattenData) GetSpecialRounds() int32 {
	if m != nil {
		return m.SpecialRounds
	}
	return 0
}

func (m *PattenData) GetPattenStr() []int32 {
	if m != nil {
		return m.PattenStr
	}
	return nil
}

func (m *PattenData) GetSpecialType() int32 {
	if m != nil {
		return m.SpecialType
	}
	return 0
}

type ResultData struct {
	ResultList []*PattenData `protobuf:"bytes,1,rep,name=resultList" json:"resultList,omitempty"`
}

func (m *ResultData) Reset()                    { *m = ResultData{} }
func (m *ResultData) String() string            { return proto.CompactTextString(m) }
func (*ResultData) ProtoMessage()               {}
func (*ResultData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ResultData) GetResultList() []*PattenData {
	if m != nil {
		return m.ResultList
	}
	return nil
}

type OneArmPlayReq struct {
	LineCnt   int32 `protobuf:"varint,1,opt,name=lineCnt" json:"lineCnt,omitempty"`
	Betting   int32 `protobuf:"varint,2,opt,name=betting" json:"betting,omitempty"`
	GameLevel int32 `protobuf:"varint,3,opt,name=gameLevel" json:"gameLevel,omitempty"`
}

func (m *OneArmPlayReq) Reset()                    { *m = OneArmPlayReq{} }
func (m *OneArmPlayReq) String() string            { return proto.CompactTextString(m) }
func (*OneArmPlayReq) ProtoMessage()               {}
func (*OneArmPlayReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *OneArmPlayReq) GetLineCnt() int32 {
	if m != nil {
		return m.LineCnt
	}
	return 0
}

func (m *OneArmPlayReq) GetBetting() int32 {
	if m != nil {
		return m.Betting
	}
	return 0
}

func (m *OneArmPlayReq) GetGameLevel() int32 {
	if m != nil {
		return m.GameLevel
	}
	return 0
}

type OneArmPlayRes struct {
	ShowTable     *PlayTableData `protobuf:"bytes,1,opt,name=showTable" json:"showTable,omitempty"`
	ResultDat     *ResultData    `protobuf:"bytes,2,opt,name=resultDat" json:"resultDat,omitempty"`
	MainGameScore int32          `protobuf:"varint,3,opt,name=mainGameScore" json:"mainGameScore,omitempty"`
	// 下面为marryGame數據
	MarryGameRounds int32           `protobuf:"varint,4,opt,name=marryGameRounds" json:"marryGameRounds,omitempty"`
	MarryGameResult *MarryGameInfos `protobuf:"bytes,5,opt,name=marryGameResult" json:"marryGameResult,omitempty"`
	MarryGameTimes  int32           `protobuf:"varint,6,opt,name=marryGameTimes" json:"marryGameTimes,omitempty"`
	TotalScore      int32           `protobuf:"varint,7,opt,name=totalScore" json:"totalScore,omitempty"`
}

func (m *OneArmPlayRes) Reset()                    { *m = OneArmPlayRes{} }
func (m *OneArmPlayRes) String() string            { return proto.CompactTextString(m) }
func (*OneArmPlayRes) ProtoMessage()               {}
func (*OneArmPlayRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *OneArmPlayRes) GetShowTable() *PlayTableData {
	if m != nil {
		return m.ShowTable
	}
	return nil
}

func (m *OneArmPlayRes) GetResultDat() *ResultData {
	if m != nil {
		return m.ResultDat
	}
	return nil
}

func (m *OneArmPlayRes) GetMainGameScore() int32 {
	if m != nil {
		return m.MainGameScore
	}
	return 0
}

func (m *OneArmPlayRes) GetMarryGameRounds() int32 {
	if m != nil {
		return m.MarryGameRounds
	}
	return 0
}

func (m *OneArmPlayRes) GetMarryGameResult() *MarryGameInfos {
	if m != nil {
		return m.MarryGameResult
	}
	return nil
}

func (m *OneArmPlayRes) GetMarryGameTimes() int32 {
	if m != nil {
		return m.MarryGameTimes
	}
	return 0
}

func (m *OneArmPlayRes) GetTotalScore() int32 {
	if m != nil {
		return m.TotalScore
	}
	return 0
}

type CenterItems struct {
	Items     []int32 `protobuf:"varint,1,rep,packed,name=items" json:"items,omitempty"`
	MultParam int32   `protobuf:"varint,3,opt,name=multParam" json:"multParam,omitempty"`
}

func (m *CenterItems) Reset()                    { *m = CenterItems{} }
func (m *CenterItems) String() string            { return proto.CompactTextString(m) }
func (*CenterItems) ProtoMessage()               {}
func (*CenterItems) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *CenterItems) GetItems() []int32 {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CenterItems) GetMultParam() int32 {
	if m != nil {
		return m.MultParam
	}
	return 0
}

type HitItems struct {
	HitItems int32 `protobuf:"varint,1,opt,name=hitItems" json:"hitItems,omitempty"`
	HitBonus int32 `protobuf:"varint,2,opt,name=hitBonus" json:"hitBonus,omitempty"`
}

func (m *HitItems) Reset()                    { *m = HitItems{} }
func (m *HitItems) String() string            { return proto.CompactTextString(m) }
func (*HitItems) ProtoMessage()               {}
func (*HitItems) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *HitItems) GetHitItems() int32 {
	if m != nil {
		return m.HitItems
	}
	return 0
}

func (m *HitItems) GetHitBonus() int32 {
	if m != nil {
		return m.HitBonus
	}
	return 0
}

type SingleTurn struct {
	CenterItems *CenterItems `protobuf:"bytes,1,opt,name=centerItems" json:"centerItems,omitempty"`
	HitItems    *HitItems    `protobuf:"bytes,2,opt,name=hitItems" json:"hitItems,omitempty"`
}

func (m *SingleTurn) Reset()                    { *m = SingleTurn{} }
func (m *SingleTurn) String() string            { return proto.CompactTextString(m) }
func (*SingleTurn) ProtoMessage()               {}
func (*SingleTurn) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *SingleTurn) GetCenterItems() *CenterItems {
	if m != nil {
		return m.CenterItems
	}
	return nil
}

func (m *SingleTurn) GetHitItems() *HitItems {
	if m != nil {
		return m.HitItems
	}
	return nil
}

type OneRound struct {
	Turns []*SingleTurn `protobuf:"bytes,1,rep,name=turns" json:"turns,omitempty"`
}

func (m *OneRound) Reset()                    { *m = OneRound{} }
func (m *OneRound) String() string            { return proto.CompactTextString(m) }
func (*OneRound) ProtoMessage()               {}
func (*OneRound) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *OneRound) GetTurns() []*SingleTurn {
	if m != nil {
		return m.Turns
	}
	return nil
}

type MarryGameInfos struct {
	ResultRounds []*OneRound `protobuf:"bytes,3,rep,name=resultRounds" json:"resultRounds,omitempty"`
}

func (m *MarryGameInfos) Reset()                    { *m = MarryGameInfos{} }
func (m *MarryGameInfos) String() string            { return proto.CompactTextString(m) }
func (*MarryGameInfos) ProtoMessage()               {}
func (*MarryGameInfos) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *MarryGameInfos) GetResultRounds() []*OneRound {
	if m != nil {
		return m.ResultRounds
	}
	return nil
}

type GMGameReq struct {
	GmCmd string `protobuf:"bytes,1,opt,name=gmCmd" json:"gmCmd,omitempty"`
}

func (m *GMGameReq) Reset()                    { *m = GMGameReq{} }
func (m *GMGameReq) String() string            { return proto.CompactTextString(m) }
func (*GMGameReq) ProtoMessage()               {}
func (*GMGameReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *GMGameReq) GetGmCmd() string {
	if m != nil {
		return m.GmCmd
	}
	return ""
}

type GMGameRes struct {
}

func (m *GMGameRes) Reset()                    { *m = GMGameRes{} }
func (m *GMGameRes) String() string            { return proto.CompactTextString(m) }
func (*GMGameRes) ProtoMessage()               {}
func (*GMGameRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func init() {
	proto.RegisterType((*PlayTableData)(nil), "msg.PlayTableData")
	proto.RegisterType((*PattenData)(nil), "msg.PattenData")
	proto.RegisterType((*ResultData)(nil), "msg.ResultData")
	proto.RegisterType((*OneArmPlayReq)(nil), "msg.OneArmPlayReq")
	proto.RegisterType((*OneArmPlayRes)(nil), "msg.OneArmPlayRes")
	proto.RegisterType((*CenterItems)(nil), "msg.CenterItems")
	proto.RegisterType((*HitItems)(nil), "msg.HitItems")
	proto.RegisterType((*SingleTurn)(nil), "msg.SingleTurn")
	proto.RegisterType((*OneRound)(nil), "msg.OneRound")
	proto.RegisterType((*MarryGameInfos)(nil), "msg.MarryGameInfos")
	proto.RegisterType((*GMGameReq)(nil), "msg.GMGameReq")
	proto.RegisterType((*GMGameRes)(nil), "msg.GMGameRes")
}

func init() { proto.RegisterFile("Game.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x57, 0xda, 0xb5, 0x27, 0x74, 0x43, 0x86, 0x8b, 0x08, 0x21, 0x34, 0x2c, 0x40, 0xe3,
	0x62, 0x85, 0x95, 0xeb, 0x5d, 0x6c, 0x45, 0x1a, 0x95, 0x36, 0x51, 0xb9, 0x7d, 0x01, 0xb7, 0x33,
	0x5d, 0x44, 0xe2, 0x74, 0xb6, 0x0b, 0xea, 0x13, 0xf2, 0x2a, 0x3c, 0x06, 0xf6, 0xb1, 0x93, 0x38,
	0xbd, 0xeb, 0xf9, 0xce, 0xe7, 0xf3, 0xf7, 0x7d, 0x0d, 0xc0, 0x2d, 0x2f, 0xc4, 0x78, 0xab, 0x4a,
	0x53, 0x92, 0x6e, 0xa1, 0x37, 0xf4, 0x02, 0x46, 0xf3, 0x9c, 0xef, 0x97, 0x7c, 0x95, 0x8b, 0x6f,
	0xdc, 0x70, 0xf2, 0x06, 0x86, 0xc6, 0x05, 0x77, 0x99, 0x36, 0x69, 0xe7, 0xac, 0x7b, 0xde, 0x63,
	0x0d, 0x40, 0xff, 0x75, 0x00, 0xe6, 0xdc, 0x18, 0x21, 0x91, 0xfc, 0x16, 0x40, 0xed, 0x5c, 0x4a,
	0x8a, 0xd9, 0x83, 0x65, 0x77, 0x2c, 0x3b, 0x42, 0xc8, 0x6b, 0x18, 0x3c, 0x66, 0x66, 0x5a, 0xee,
	0xa4, 0x49, 0x8f, 0x30, 0x5b, 0xc7, 0xae, 0x91, 0xfd, 0x3d, 0x33, 0xa2, 0xb0, 0x4f, 0xbb, 0x98,
	0x6c, 0x80, 0xf0, 0x72, 0x99, 0x15, 0x42, 0xa7, 0xcf, 0xea, 0x97, 0x18, 0x93, 0xf7, 0x30, 0xd2,
	0x5b, 0xb1, 0xce, 0x78, 0xce, 0x6c, 0xa5, 0x07, 0x9d, 0xf6, 0x90, 0xd0, 0x06, 0x5d, 0xfd, 0x2d,
	0x4e, 0xba, 0x30, 0x2a, 0xed, 0xfb, 0x45, 0x6a, 0x80, 0x9c, 0x41, 0x12, 0xe8, 0xcb, 0xfd, 0x56,
	0xa4, 0xc7, 0x58, 0x21, 0x86, 0xe8, 0x15, 0x00, 0x13, 0x7a, 0x97, 0x1b, 0xdc, 0xf4, 0xb3, 0xdd,
	0x14, 0xa3, 0xfa, 0x2e, 0xc9, 0xe4, 0x74, 0x6c, 0x2f, 0x38, 0x6e, 0xce, 0xc1, 0x22, 0x0a, 0xe5,
	0x30, 0xfa, 0x21, 0xc5, 0xb5, 0x2a, 0xdc, 0x79, 0x99, 0x78, 0x22, 0x29, 0x1c, 0xe7, 0xf6, 0x2a,
	0x53, 0x69, 0xc2, 0xa1, 0xaa, 0xd0, 0x65, 0x56, 0xc2, 0x98, 0x4c, 0x6e, 0xc2, 0x91, 0xaa, 0xd0,
	0xed, 0xb0, 0xb1, 0x82, 0xdd, 0x89, 0xdf, 0x22, 0xaf, 0x6e, 0x54, 0x03, 0xf4, 0xef, 0x51, 0xbb,
	0x87, 0x26, 0x5f, 0x60, 0xa8, 0x1f, 0xcb, 0x3f, 0xa8, 0x26, 0x76, 0x49, 0x26, 0xc4, 0x0f, 0x19,
	0x6b, 0xcc, 0x1a, 0x12, 0xb9, 0x80, 0xa1, 0xaa, 0xb6, 0xc4, 0xee, 0xd5, 0x5a, 0xcd, 0xee, 0xac,
	0x61, 0xb8, 0xd3, 0x17, 0x3c, 0x93, 0xce, 0x45, 0x8b, 0x75, 0xa9, 0x44, 0x18, 0xaa, 0x0d, 0x92,
	0x73, 0x38, 0x2d, 0xb8, 0x52, 0x7b, 0x87, 0x04, 0x89, 0xbc, 0x86, 0x87, 0x30, 0xb9, 0x8a, 0x99,
	0xd8, 0x05, 0xc5, 0x4c, 0x26, 0x2f, 0x71, 0x88, 0xfb, 0x2a, 0x37, 0x93, 0x3f, 0x4b, 0xcd, 0x0e,
	0xb9, 0xe4, 0x23, 0x9c, 0xd4, 0x90, 0xf7, 0x4a, 0x1f, 0xfb, 0x1c, 0xa0, 0xce, 0xa7, 0xa6, 0x34,
	0x3c, 0xf7, 0x33, 0x7b, 0xb1, 0x23, 0x84, 0x5e, 0x43, 0x32, 0x15, 0xd2, 0x08, 0xe5, 0xdc, 0xa7,
	0xc9, 0x2b, 0xe8, 0x65, 0xee, 0x47, 0xf0, 0xbf, 0x0f, 0x9c, 0x18, 0x85, 0x6d, 0x3a, 0xe7, 0x8a,
	0x17, 0x95, 0x18, 0x35, 0x40, 0x6f, 0x60, 0xf0, 0xdd, 0xbb, 0x57, 0x07, 0xf3, 0xce, 0x42, 0x89,
	0xca, 0xbc, 0x71, 0xee, 0xa6, 0x94, 0x3b, 0x1d, 0xfd, 0x25, 0x30, 0xa6, 0xbf, 0x00, 0x16, 0x56,
	0xf6, 0x5c, 0x2c, 0x77, 0x4a, 0x92, 0x09, 0x24, 0xeb, 0x66, 0xa8, 0x20, 0xe7, 0x0b, 0xbc, 0x4b,
	0x34, 0x2c, 0x8b, 0x49, 0xe4, 0x53, 0xd4, 0xd9, 0xab, 0x39, 0xc2, 0x07, 0xd5, 0x68, 0xcd, 0x20,
	0xf4, 0x12, 0x06, 0xd6, 0x3c, 0xa8, 0x03, 0xf9, 0x00, 0x3d, 0x63, 0x5b, 0xea, 0x96, 0xb1, 0x9b,
	0x51, 0x98, 0xcf, 0xd2, 0x29, 0x9c, 0xb4, 0x15, 0x21, 0x97, 0xf0, 0xdc, 0x9b, 0x23, 0xc8, 0xdc,
	0xc5, 0xf7, 0xbe, 0x67, 0x55, 0x9d, 0xb5, 0x28, 0xf4, 0x1d, 0x0c, 0x6f, 0xef, 0xbd, 0x86, 0x4f,
	0xee, 0xd2, 0x9b, 0x62, 0x5a, 0xf8, 0x6f, 0xc7, 0x90, 0xf9, 0x80, 0x26, 0x0d, 0x45, 0xaf, 0xfa,
	0xf8, 0xb5, 0xfa, 0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x4d, 0x57, 0x02, 0xbb, 0x04, 0x00,
	0x00,
}