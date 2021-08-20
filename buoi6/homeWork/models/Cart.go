package models

type Cart struct {
	CustomID   string
	ID         string `gorm:"Primary_key"`
	ProductsID string
	Numbers    string
	Status     string
}
