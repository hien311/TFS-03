package db

import (
	"todo_mvc/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func Connect() *gorm.DB {
	evn, err := util.LoadDBConfig("./util")
	if err != nil {
		logrus.Error(err)
	}

	db, err := gorm.Open(evn.DBDriver, evn.DSN)
	if err != nil {
		logrus.Error("Fail to connect DB ", err)
	}
	return db
}
