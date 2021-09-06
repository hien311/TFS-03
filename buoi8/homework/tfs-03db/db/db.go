package db

import (
	"tfs-03db/model"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:giadinh310@tcp(localhost:3306)/tfs-03")
	if err != nil {
		logrus.Error("loi ket noi db", err)
	}
	return db
}

func CreateTable() {
	database := Connect()
	var people model.People
	database.DropTableIfExists("peoples")
	err := database.AutoMigrate(&people)
	if err != nil {
		logrus.Error(err)
	}
}
