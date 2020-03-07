package common

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisConnInit(t *testing.T) {
	godotenv.Load(".env.example")
	RedisConnInit()
	redisClient := RedisConnManager.Get("default").GetRedisClient()
	assert.True(t, redisClient != nil)
}
