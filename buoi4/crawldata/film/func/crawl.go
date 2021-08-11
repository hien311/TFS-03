package crawl

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type film struct {
	name     string
	year     string
	director string
	writers  string
	link     string
	rating   string
}

func CrawlData(url string) map[int]film {
	var films = make(map[int]film)
	u, _ := http.Get(url)
	data, _ := goquery.NewDocumentFromReader(u.Body)
	data.Find(".lister-list tr").Each(func(i int, s *goquery.Selection) {
		// bien tam de  them du lieu vao
		temp := film{}

		temp.rating, _ = s.Find(".posterColumn [name = ir]").Attr("data-value")
		temp.year = s.Find(".titleColumn span").Text()
		info := s.Find(".titleColumn a")
		temp.name = info.Text()
		temp.link, _ = info.Attr("href")
		temp.link = "https://www.imdb.com/" + temp.link
		title, _ := info.Attr("title")
		titleSplit := strings.Split(title, ", ")
		temp.director = titleSplit[0]
		temp.writers = strings.Join(titleSplit[1:], ", ")
		films[i+1] = temp
	})
	return films
}
