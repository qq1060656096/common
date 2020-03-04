package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	godotenv.Load(".env")
	fmt.Printf("%v", DbConnManager.length())
	assert.Equal(t, true, false)
}