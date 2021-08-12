package crawl

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CrawlData(url string) map[int]Film {
	var films = make(map[int]Film)
	res, _ := http.Get(url)
	data, _ := goquery.NewDocumentFromReader(res.Body)
	data.Find(".lister-list tr").Each(func(i int, s *goquery.Selection) {
		// bien tam de  them du lieu vao
		temp := Film{}
		temp.Rank = i
		temp.Rating, _ = s.Find(".posterColumn [name = ir]").Attr("data-value")
		temp.Year = s.Find(".titleColumn span").Text()
		info := s.Find(".titleColumn a")
		temp.Name = info.Text()
		temp.Link, _ = info.Attr("href")
		temp.Link = "https://www.imdb.com/" + temp.Link
		title, _ := info.Attr("title")
		titleSplit := strings.Split(title, ", ")
		temp.Director = titleSplit[0]
		temp.Writers = strings.Join(titleSplit[1:], ", ")
		films[i+1] = temp
	})
	return films
}
