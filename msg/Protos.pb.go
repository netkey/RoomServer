// Code generated by protoc-gen-go.
// source: Protos.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MsgID int32

const (
	MsgID_None              MsgID = 0
	MsgID_loginReq          MsgID = 1000
	MsgID_loginRes          MsgID = 1001
	MsgID_playerInfoUpdate  MsgID = 1002
	MsgID_getVipShopListReq MsgID = 1100
	MsgID_getVipShopListRes MsgID = 1101
	// DeskRelate
	MsgID_createDeskReq  MsgID = 2001
	MsgID_createDeskRes  MsgID = 2002
	MsgID_checkInDeskReq MsgID = 2003
	MsgID_checkInDeskRes MsgID = 2004
	MsgID_gameDataInfo   MsgID = 2005
	MsgID_gmCmdReq       MsgID = 18000
	MsgID_gmCmdRes       MsgID = 18001
	MsgID_heartBeat      MsgID = 9999
)

var MsgID_name = map[int32]string{
	0:     "None",
	1000:  "loginReq",
	1001:  "loginRes",
	1002:  "playerInfoUpdate",
	1100:  "getVipShopListReq",
	1101:  "getVipShopListRes",
	2001:  "createDeskReq",
	2002:  "createDeskRes",
	2003:  "checkInDeskReq",
	2004:  "checkInDeskRes",
	2005:  "gameDataInfo",
	18000: "gmCmdReq",
	18001: "gmCmdRes",
	9999:  "heartBeat",
}
var MsgID_value = map[string]int32{
	"None":              0,
	"loginReq":          1000,
	"loginRes":          1001,
	"playerInfoUpdate":  1002,
	"getVipShopListReq": 1100,
	"getVipShopListRes": 1101,
	"createDeskReq":     2001,
	"createDeskRes":     2002,
	"checkInDeskReq":    2003,
	"checkInDeskRes":    2004,
	"gameDataInfo":      2005,
	"gmCmdReq":          18000,
	"gmCmdRes":          18001,
	"heartBeat":         9999,
}

func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}
func (MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type LoginTypes int32

const (
	LoginTypes_Normal_Login LoginTypes = 0
	LoginTypes_Face_Book    LoginTypes = 1
	LoginTypes_ReConnect    LoginTypes = 2
)

var LoginTypes_name = map[int32]string{
	0: "Normal_Login",
	1: "Face_Book",
	2: "ReConnect",
}
var LoginTypes_value = map[string]int32{
	"Normal_Login": 0,
	"Face_Book":    1,
	"ReConnect":    2,
}

func (x LoginTypes) String() string {
	return proto.EnumName(LoginTypes_name, int32(x))
}
func (LoginTypes) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func init() {
	proto.RegisterEnum("msg.MsgID", MsgID_name, MsgID_value)
	proto.RegisterEnum("msg.LoginTypes", LoginTypes_name, LoginTypes_value)
}

func init() { proto.RegisterFile("Protos.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0xbf, 0xf6, 0x53, 0x9b, 0x1e, 0xda, 0xe6, 0x74, 0x44, 0x2f, 0xa2, 0x0b, 0x37, 0x6e,
	0x5d, 0xb5, 0x41, 0x08, 0xd4, 0x20, 0xf5, 0x67, 0x5b, 0xc6, 0xf4, 0x38, 0x09, 0x49, 0x66, 0xc6,
	0x9c, 0xd9, 0xf4, 0x1a, 0x14, 0xbc, 0x31, 0x05, 0xeb, 0xcf, 0x3d, 0xa8, 0x57, 0x61, 0x27, 0x20,
	0x14, 0x75, 0x35, 0x3c, 0x0f, 0xef, 0xcb, 0xcc, 0xbc, 0xd0, 0x3b, 0xad, 0x8d, 0x33, 0x7c, 0x60,
	0xfd, 0x21, 0xfe, 0x57, 0xac, 0x46, 0xb7, 0x6d, 0xd8, 0x3e, 0x61, 0x15, 0x47, 0x22, 0x80, 0xad,
	0xc4, 0x68, 0xc2, 0x7f, 0xa2, 0x0f, 0x41, 0x69, 0x54, 0xae, 0x67, 0x74, 0x83, 0xef, 0x9d, 0x0d,
	0x64, 0xfc, 0xe8, 0x88, 0x3d, 0x40, 0x5b, 0xca, 0x25, 0xd5, 0xb1, 0xbe, 0x36, 0x17, 0x76, 0x21,
	0x1d, 0xe1, 0x67, 0x47, 0xec, 0xc3, 0x50, 0x91, 0xbb, 0xcc, 0xed, 0x59, 0x66, 0xec, 0x34, 0x67,
	0xe7, 0xdb, 0x0f, 0xc1, 0x5f, 0x9e, 0xf1, 0x31, 0x10, 0x02, 0xfa, 0x69, 0x4d, 0xeb, 0x72, 0x44,
	0x5c, 0xf8, 0xec, 0x2a, 0xfc, 0xe9, 0x18, 0x9f, 0x43, 0xb1, 0x0b, 0x83, 0x34, 0xa3, 0xb4, 0x88,
	0xf5, 0x77, 0xf0, 0xe5, 0xb7, 0x64, 0x7c, 0x0d, 0xc5, 0x10, 0x7a, 0x4a, 0x56, 0x14, 0x49, 0x27,
	0xfd, 0xd3, 0xf0, 0x2d, 0x14, 0x03, 0x08, 0x54, 0x35, 0xa9, 0x16, 0xbe, 0xf6, 0x74, 0xd7, 0xda,
	0x60, 0xc6, 0x55, 0xc3, 0xdd, 0x8c, 0x64, 0xed, 0xc6, 0xeb, 0x4b, 0xf1, 0x3e, 0x19, 0x1d, 0x01,
	0x4c, 0xfd, 0x57, 0xcf, 0x97, 0x96, 0x58, 0x20, 0xf4, 0x12, 0x53, 0x57, 0xb2, 0x9c, 0x37, 0xb2,
	0x59, 0xa6, 0x7b, 0x2c, 0x53, 0x9a, 0x8f, 0x8d, 0x29, 0xb0, 0xe5, 0x71, 0x46, 0x13, 0xa3, 0x35,
	0xa5, 0x0e, 0xdb, 0x57, 0x3b, 0xcd, 0xae, 0x87, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x7f,
	0xce, 0x40, 0x67, 0x01, 0x00, 0x00,
}
