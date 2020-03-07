package common

import (
	"fmt"
	gormManager "github.com/qq1060656096/go-gorm-manager"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	DefaultDbConnName  = "default"
	CommonDbConnName   = "common"
	BusinessDbConnName = "business"
)

var DbConnManager = gormManager.NewConnectionManager()

func DbConnInit() {
	db, err := OsEnvManager.GetBool("DB")
	if err != nil {
		logrus.Infof("common.DbConnInit.os.env.key.DB:%s", err)
	}

	if db {
		DefaultDbConnInit()
	}

	common, err := OsEnvManager.GetBool("COMMON_DB")
	if err != nil {
		logrus.Infof("common.DbConnInit.os.env.key.COMMON_DB:%s", err)
	}
	if common {
		CommonDbConnInit()
	}

	business, err := OsEnvManager.GetBool("BUSINESS_DB")
	if err != nil {
		logrus.Infof("common.DbConnInit.os.env.key.BUSINESS_DB:%s", err)
	}
	if business {
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
