package db

import (
	crawl "crawl/func"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, data crawl.Films) {
	db.AutoMigrate(&crawl.Film{})

	for i := 0; i < len(data); i++ {
		db.Create(&crawl.Film{Rank: data[i].Rank, Name: data[i].Name, Year: data[i].Year, Rating: data[i].Rating, Director: data[i].Director, Link: data[i].Link, Writers: data[i].Writers})
	}

	/*
		done := make(chan bool)
		ch := make(chan crawl.Film, 10)
		go func() {
			for i := 1; i <= len(data); i++ {
				ch <- data[i]
			}
			close(ch)
			done <- true
		}()
		for v := range ch {
			db.Create(&crawl.Film{Rank: v.Rank, Name: v.Name, Year: v.Year, Rating: v.Rating, Director: v.Director, Link: v.Link, Writers: v.Writers})
		}
		<-done */ //Hình như đoạn này dùng Goroutine chậm hơn đoạn code bên trên đúng ko ạ?

}

func GetAll(db *gorm.DB) crawl.Films {
	var Films crawl.Films
	db.Find(&Films)
	return Films
}
