package crawl

type Film struct {
	Rank     int    `json:"rank"`
	Name     string `json:"name"`
	Year     string `json:"year"`
	Director string `json:"director"`
	Writers  string `json:"writers"`
	Link     string `json:"link"`
	Rating   string `json:"Rating"`
}
