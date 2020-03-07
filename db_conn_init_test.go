package common

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	godotenv.Load(".env.example")
	DbConnInit()
	db, err := DbConnManager.Get("common").GetGormDB()
	assert.True(t, err == nil)
	assert.True(t, db != nil)
	assert.Equal(t, 3, DbConnManager.Length())
}
