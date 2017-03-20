package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"RoomServer/conf"

)

var (
	RAWDB *sql.DB
)

type MYDB struct {
	*sql.DB
	SQLStr string
}

func (this *MYDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return this.DB.Query(query, args)
}

func MySQLEscape(v string) string {
	var pos = 0
	buf := make([]byte, 2*len(v))
	for i := 0; i < len(v); i++ {
		c := v[i]
		switch c {
		case '\x00':
			buf[pos] = '\\'
			buf[pos+1] = '0'
			pos += 2
		case '\n':
			buf[pos] = '\\'
			buf[pos+1] = 'n'
			pos += 2
		case '\r':
			buf[pos] = '\\'
			buf[pos+1] = 'r'
			pos += 2
		case '\x1a':
			buf[pos] = '\\'
			buf[pos+1] = 'Z'
			pos += 2
		case '\'':
			buf[pos] = '\\'
			buf[pos+1] = '\''
			pos += 2
		case '"':
			buf[pos] = '\\'
			buf[pos+1] = '"'
			pos += 2
		case '\\':
			buf[pos] = '\\'
			buf[pos+1] = '\\'
			pos += 2
		default:
			buf[pos] = c
			pos++
		}
	}
	return string(buf[:pos])
}

func SetupDBMysql() {
	var DBPassword string
	var err error
	DBPassword = conf.Server.MysqlPassWord


	str := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local&autocommit=true&charset=utf8mb4",
		conf.Server.MysqlUserName,
		DBPassword,
		conf.Server.MysqlServer,
		conf.Server.MysqlPort,
		conf.Server.MysqlName,
	)

	db2, err := sql.Open("mysql", str)
	if err != nil {
		panic(err)
	}
	RAWDB = db2
	RAWDB.SetMaxOpenConns(100)
	RAWDB.SetMaxIdleConns(10)
}
