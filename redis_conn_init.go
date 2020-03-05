package common

import (
	redisManager "github.com/qq1060656096/go-redis-manager"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	DefaultRedisConnName = "default"
	AuthRedisConnName = "auth"
)

var RedisConnManager = redisManager.NewConnectionManager()


func RedisConnInit()  {
	dr, err := OsEnvManager.GetBool("DEFAULT_REDIS")
	if err != nil {
		logrus.Infof("common.RedisConnInit.os.env.key.DEFAULT_REDIS:%s", err)
	}
	if dr {
		DefaultRedisConnInit()
	}

	ar, err := OsEnvManager.GetBool("AUTH_REDIS")
	if err != nil {
		logrus.Infof("common.RedisConnInit.os.env.key.AUTH_REDIS:%s", err)
	}
	if ar {
		AuthRedisConnInit()
	}
}

func DefaultRedisConnInit()  {
	addr := os.Getenv("DEFAULT_REDIS_ADDR")
	pass := os.Getenv("DEFAULT_REDIS_PASSWORD")
	db, err := OsEnvManager.GetInt("DEFAULT_REDIS_DB")
	if err != nil {
		logrus.Infof("common.DefaultRedisConnInit.os.env.key.DEFAULT_REDIS_DB:%s", err)
	}
	ConnectionRedisInit(DefaultRedisConnName, &redis.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       db,        // use default DB
	})
}


func AuthRedisConnInit()  {
	addr := os.Getenv("AUTH_REDIS_ADDR")
	pass := os.Getenv("AUTH_REDIS_PASSWORD")
	db, err := OsEnvManager.GetInt("AUTH_REDIS_DB")
	if err != nil {
		logrus.Infof("common.AuthRedisConnInit.os.env.key.AUTH_REDIS_DB:%s", err)
	}
	ConnectionRedisInit(AuthRedisConnName, &redis.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       db,        // use default DB
	})
}

func ConnectionRedisInit(connName string, options *redis.Options)  {
	RedisConnManager.Add(connName, options)
}