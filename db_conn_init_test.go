package common

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	godotenv.Load("config/.db.env")
	DbConnInit()
	db, err := GetDefaultDbConn()
	db, err = GetCommonDbConn()
	db, err = GetBusinessDbConn()

	assert.True(t, err == nil)
	assert.True(t, db != nil)
	assert.Equal(t, 3, DbConnManager.Length())
}
