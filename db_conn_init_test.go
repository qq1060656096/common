package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


func TestInit(t *testing.T) {
	godotenv.Load(".env.example")
	DbConnInit()
	db, err := DbConnManager.Get("common").GetGormDB()
	fmt.Printf("%v\n%#v", db, err)

	assert.Equal(t, 3, DbConnManager.Length())
}