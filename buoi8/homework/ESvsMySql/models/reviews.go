package model

type Review struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Body   string `json:"body" gorm:"type:longtext"`
}
