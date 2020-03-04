package common

import (
	"fmt"
	"github.com/go-redis/redis"
	gormManager "github.com/qq1060656096/go-gorm-manager"
	redisManager "github.com/qq1060656096/go-redis-manager"
	"os"
	"strconv"
)

var DbConnManager = gormManager.NewConnectionManager()
var RedisConnManager = redisManager.NewConnectionManager()

const (
	DefaultDbConnName  = "default"
	CommonDbConnName   = "common"
	BusinessDbConnName = "business"

	DefaultRedisConnName = "default"
	AuthRedisConnName = "auth"
)

func init()  {
	DbConnInit()
	RedisInit()
}

func DbConnInit()  {
	fmt.Println(os.Getenv("DEFAULT_DB"))
	if os.Getenv("DEFAULT_DB") == "true" {
		DefaultDbConnInit()
	}
	if os.Getenv("COMMON_DB") == "true" {
		CommonDbConnInit()
	}
	if os.Getenv("BUSINESS_DB") == "true" {
		BusinessDbConnInit()
	}
}
func DefaultDbConnInit() {
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	charset := os.Getenv("DB_CHARSET")
	ConnectDbMySQL(DefaultDbConnName, user, pass, host, port, database, charset)
}

func CommonDbConnInit() {
	user := os.Getenv("COMMON_DB_USERNAME")
	pass := os.Getenv("COMMON_DB_PASSWORD")
	host := os.Getenv("COMMON_DB_HOST")
	port := os.Getenv("COMMON_DB_PORT")
	database := os.Getenv("COMMON_DB_DATABASE")
	charset := os.Getenv("COMMON_DB_CHARSET")
	ConnectDbMySQL(CommonDbConnName, user, pass, host, port, database, charset)
}

func BusinessDbConnInit() {
	user := os.Getenv("BUSINESS_DB_USERNAME")
	pass := os.Getenv("BUSINESS_DB_PASSWORD")
	host := os.Getenv("BUSINESS_DB_HOST")
	port := os.Getenv("BUSINESS_DB_PORT")
	database := os.Getenv("BUSINESS_DB_DATABASE")
	charset := os.Getenv("BUSINESS_DB_CHARSET")
	ConnectDbMySQL(BusinessDbConnName, user, pass, host, port, database, charset)
}

func ConnectDbMySQL(connName, user, pass, host, port, database, charset string) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		database,
		charset,
	)
	DbConnManager.Add(connName, &gormManager.ConnectionConfig{
		DatabaseDriverName: gormManager.DRIVER_MY_SQL,
		DataSourceName:     dataSourceName,
	})
}

func RedisInit()  {
	if os.Getenv("AUTH_REDIS") == "true" {
		AuthRedisInit()
	}
	if os.Getenv("DEFUALT_REDUS") == "true" {
		DefaultRedisInit()
	}
}

func AuthRedisInit()  {
	addr := os.Getenv("AUTH_REDIS_ADDR")
	pass := os.Getenv("AUTH_REDIS_PASSWORD")
	db := os.Getenv("AUTH_REDIS_DB")
	dbInt, _ := strconv.Atoi(db)
	ConnectionRedisInit(DefaultRedisConnName, &redis.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       dbInt,        // use default DB
	})
}

func DefaultRedisInit()  {
	addr := os.Getenv("DEFAULT_REDIS_ADDR")
	pass := os.Getenv("DEFAULT_REDIS_PASSWORD")
	db := os.Getenv("DEFAULT_REDIS_DB")
	dbInt, _ := strconv.Atoi(db)
	ConnectionRedisInit(DefaultRedisConnName, &redis.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       dbInt,        // use default DB
	})
}

func ConnectionRedisInit(connName string, options *redis.Options)  {
	redisManager.Add(connName, options)
}