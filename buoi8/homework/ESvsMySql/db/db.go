package db

import (
	model "buoi8_homework/models"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var (
	review model.Review
	ch     = make(chan model.Review)
	mutex  sync.Mutex
	wg     sync.WaitGroup
	ctx    = context.Background()
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:giadinh310@tcp(localhost:3306)/tfs-03")
	if err != nil {
		logrus.Info(err)
	}
	return db
}

func UpToES() {
	id := 0
	es, err := elastic.NewClient()
	if err != nil {
		logrus.Error("loi o khoi tao: ", err)
	}

	csvData, err := os.Open("./data/train.csv")
	if err != nil {
		logrus.Error("loi doc file: ", err)
	}
	defer csvData.Close()
	wg.Add(1)
	train := csv.NewReader(csvData)
	for j := 1; j < 101; j++ {
		go func() {
			for {
				mutex.Lock()
				line, err := read(train)
				if err == io.EOF {
					logrus.Info("close chan ", id)
					close(ch)
					break
				}
				if err != nil {
					logrus.Error("loi ghi")
				}
				review = model.Review{
					Status: line[0],
					Title:  line[1],
					Body:   line[2],
				}
				ch <- review
				mutex.Unlock()
			}
		}()
	}
	for i := 0; i < 130; i++ {
		go func() {
			for {
				item, ok := <-ch
				if !ok {
					logrus.Info("done")
					wg.Done()
					return
				}
				t, err := json.Marshal(item)
				if err != nil {
					logrus.Error(err)
				}
				es.Index().Index("reviews").BodyJson(string(t)).Do(ctx)

			}
		}()
	}
	wg.Wait()
}

func UpToMySqlDb() {
	database := Connect()
	defer database.Close()

	csvData, err := os.Open("./data/train.csv")
	if err != nil {
		logrus.Error(err)
	}
	defer csvData.Close()
	train := csv.NewReader(csvData)
	wg.Add(1)

	for j := 1; j < 101; j++ {
		go func() {
			for {
				mutex.Lock()
				line, err := train.Read()
				if err == io.EOF {
					logrus.Info("close chan")
					close(ch)
					break
				}
				if err != nil {
					logrus.Error("loi ghi")
				}
				review = model.Review{
					Status: line[0],
					Title:  line[1],
					Body:   line[2],
				}
				ch <- review
				mutex.Unlock()
			}
		}()
	}
	for i := 0; i < 130; i++ {
		go func() {
			for {
				item, ok := <-ch
				if !ok {
					logrus.Info("done")
					wg.Done()
					return
				}
				insert, err := database.Query("INSERT INTO reviews(status, title, body) VALUES ( ?, ?, ? )", item.Status, item.Title, item.Body)
				if err != nil {
					logrus.Error(err)
				}
				insert.Close()
			}
		}()
	}
	wg.Wait()
}

func read(file *csv.Reader) ([]string, error) {
	return file.Read()
}
