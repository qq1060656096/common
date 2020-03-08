package common

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisConnInit(t *testing.T) {
	godotenv.Load("config/.cache.env")
	RedisConnInit()
	redisClient := GetDefaultRedisConn()
	redisClient = GetAuthRedisConn()
	assert.True(t, redisClient != nil)
}
