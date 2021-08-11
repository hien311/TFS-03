package main

import (
	crawl "crawl/func"
	"fmt"
)

func main() {
	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	fmt.Println(crawl.CrawlData(url))
}
