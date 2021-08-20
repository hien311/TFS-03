package models

type User struct {
	UserName    string `gorm:"Primary_key"`
	Password    string
	PhoneNumber string
}
