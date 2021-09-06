package models

type Teacher struct {
	Teacher_ID int `gorm:"Primary_key"`
	Name       string
	Age        int
	Phone      int
	Address    string
}
