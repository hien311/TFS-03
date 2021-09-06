package main

import (
	db "crawl/database"
	crawl "crawl/func"
)

func main() {
	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	data := crawl.CrawlData(url)
	database := db.Connect()
	db.Create(database, data)
}
