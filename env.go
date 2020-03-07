package common

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var OsEnvManager = &OsEnv{}

type OsEnv struct {
}

func (e *OsEnv) Get(key string) string {
	s := os.Getenv(key)
	logrus.Infof("common.Get.%s", s)
	return s
}

func (e *OsEnv) GetBool(key string) (bool, error) {
	v, err := strconv.ParseBool(e.Get(key))
	return v, err
}

func (e *OsEnv) GetInt64(key string) (int64, error) {
	v, err := strconv.ParseInt(e.Get(key), 10, 64)
	return v, err
}

func (e *OsEnv) GetInt32(key string) (int32, error) {
	v, err := e.GetInt64(key)
	return int32(v), err
}

func (e *OsEnv) GetInt(key string) (int, error) {
	v, err := e.GetInt64(key)
	return int(v), err
}

func (e *OsEnv) GetFloat64(key string) (float64, error) {
	v, err := strconv.ParseFloat(e.Get(key), 64)
	return v, err
}

func (e *OsEnv) GetFloat32(key string) (float32, error) {
	v, err := e.GetFloat64(key)
	return float32(v), err
}
