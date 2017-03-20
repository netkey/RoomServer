package db

import (
	"time"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"RoomServer/conf"
	"github.com/name5566/leaf/log"
)

var (
	redisPool *redis.Pool
)

func newPool(server, password string, maxidle int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxidle,
		IdleTimeout: 0,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitRedis() {
	fmt.Println(conf.Server.RedisAddr)
	log.Debug("LogRedis Add:[%s] PW:[%s] MC:[%d]", conf.Server.RedisAddr, conf.Server.RedisPasswd, conf.Server.RedisMaxConn)
	redisPool = newPool(conf.Server.RedisAddr, conf.Server.RedisPasswd, conf.Server.RedisMaxConn)
}
