package db

import (
	"tfs-03db/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

type Database struct {
	*gorm.DB
}

func Connect() *Database {
	conn, err := gorm.Open("mysql", "root:giadinh310@tcp(localhost:3306)/tfs-03")
	if err != nil {
		logrus.Error("loi ket noi db", err)
	}
	return &Database{conn}
}

func (db *Database) ServeInit() {
	var people model.People

	db.DropTableIfExists("peoples")
	db.AutoMigrate(&people)
}
