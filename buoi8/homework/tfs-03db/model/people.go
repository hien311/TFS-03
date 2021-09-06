package model

type People struct {
	ID      uint   `json:"id" gorm:"Primary_key"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Phone   int    `json:"phone"`
	Role    string `json:"role"`
	Deleted bool   `json:"deleted"`
}

type Peoples []People
