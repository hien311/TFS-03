package models

import "time"

type Lecture struct {
	Lecture_ID int `gorm:"Primary_key"`
	Name       string
	Des        string
	Time       time.Time
	Teacher_ID int
}
