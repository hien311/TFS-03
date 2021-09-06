package models

type Student struct {
	Student_ID int `gorm:"Primary_key"`
	Name       string
	Age        string
	Address    string
	Phone      int
}
