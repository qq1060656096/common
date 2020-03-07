package common

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisConnInit(t *testing.T) {
	godotenv.Load(".env.example")
	RedisConnInit()
	redisClient := GetDefaultRedisConn()
	redisClient = GetAuthRedisConn()
	assert.True(t, redisClient != nil)
}
