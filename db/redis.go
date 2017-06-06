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

func RedisKVGet(Key string)(Val string,IsExist bool,err error){
	log.Debug("[RedisKVGet] K:(%s)",Key)
	redisConn := redisPool.Get()
	defer redisConn.Close()
	r, err := redisConn.Do("exists", Key)
	IsExist, err = redis.Bool(r, err)
	if err != nil {
		log.Debug("RedisKVGet error (%s)", err.Error())
		return
	}
	if !IsExist{
		log.Debug("[RedisKVGet] Not Exist")
		return
	}
	r, err = redisConn.Do("hgetall", Key)
	retMap, err := redis.StringMap(r, err)
	if err != nil {
		log.Debug("[RedisKVGet]")
		return
	}
	Val = retMap["Val"]
	if err != nil {
		log.Debug("[RedisKVGet]")
		return
	}
	return
}


func RedisSet(key string, val string, expire int) error {
	//mylog.LOG.I("[RedisSet] K:(%s)  V:(%s)",key,val)
	redisConn := redisPool.Get()
	defer redisConn.Close()
	_, err := redisConn.Do("set", key, val)
	if err != nil {
		log.Debug("RedisSet error (%s)", key)
		return err
	}
	if expire != 0 {
		_, err = redisConn.Do("expire", key, expire)
		if err != nil {
			log.Debug("Redis expire error: %v", err)
		}
	}
	return nil
}