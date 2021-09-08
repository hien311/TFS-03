package model

type Task struct {
	ID     int    `json:"id" gorm:"Primary_key"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
