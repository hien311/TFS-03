package models

type Produce struct {
	ID        string `gorm:"Primary_key"`
	Name      string
	Des       string
	Available string
	Price     string
}

type Produces []Produce
