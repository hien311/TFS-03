package db

import (
	"homework/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	log.Info("started serve")
	dsn := "root:giadinh310@tcp(127.0.0.1:3306)/shops"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("loi ket noi db")
	}
	return db
}

func CreateTable() {
	database := Connect()
	database.Debug().Begin().AutoMigrate(&models.Produce{})
	database.AutoMigrate(&models.Cart{})
}
