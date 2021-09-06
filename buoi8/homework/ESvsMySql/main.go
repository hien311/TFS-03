package main

import (
	db "buoi8_homework/db"
	"context"
	"fmt"
	"time"

	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func main() {
	s := "nice"
	sqlTime := sqlSearchTime(s)
	esTime := esSearchTime(s)
	fmt.Printf("thoi gian tim kiem \"%s\" voi mySql: %v", s, sqlTime)
	fmt.Printf("\nthoi gian tim kiem \"%s\" voi ES: %v", s, esTime)
}

func sqlSearchTime(s string) time.Duration {
	timeStart := time.Now()
	database := db.Connect()
	string := fmt.Sprintf("%s%s%s", "%", s, "%")
	database.Query("SElECT * FROM reviews WHERE body LIKE ?", string)
	searchTime := time.Since(timeStart)
	return searchTime
}

func esSearchTime(s string) time.Duration {
	timeStart := time.Now()
	es, err := elastic.NewClient()
	if err != nil {
		logrus.Error(err)
	}
	query := elastic.NewMatchPhraseQuery("body", s)
	es.Search().Index("reviews").Query(query).Do(ctx)
	if err != nil {
		logrus.Error(err)
	}

	searchTime := time.Since(timeStart)
	return searchTime
}
