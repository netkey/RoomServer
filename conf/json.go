package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
	MysqlServer	string
	MysqlPort   string
	MysqlUserName	string
	MysqlPassWord	string
	MysqlName		string
	RedisAddr		string
	RedisPasswd		string
	RedisMaxConn	int
	IsStg			bool
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
