package main

import (
	"buoi8_onclass/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	database := db()
	database.AutoMigrate(&models.Student{})
	database.AutoMigrate(&models.Lecture{})
	database.AutoMigrate(&models.Teacher{})
	database.AutoMigrate(&models.StuLec{})
}

func db() *gorm.DB {
	dsn := "root:giadinh310@tcp(127.0.0.1:3306)/tfs-03"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Info(err)
	}
	return db
}
