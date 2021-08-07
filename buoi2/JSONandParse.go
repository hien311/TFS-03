package main

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	UserID   string `json:"userID"`
	Username string `json:"Username"`
	Password string `json:"12345678"`
	Hobby    Hobby  `json:"hobby"`
}

type Hobby struct {
	Fishing  bool
	Shopping bool
}

func main() {
	// Khai bao bien user
	userHobby := Hobby{Fishing: true, Shopping: true}
	user := Users{
		UserID:   "1",
		Username: "Golang4dummy",
		Password: "12345678",
		Hobby:    userHobby,
	}
	/// JSON with MarshalIndent
	userJSON, err := json.MarshalIndent(user, "", "   ")
	if err != nil {
		fmt.Println("co loi!!")
	}
	//ParseJSON
	var user1 Users
	json.Unmarshal([]byte(userJSON), &user1)

	// In ra man hinh
	fmt.Println("user sau khi chuyen sang dinh dang JSON =>>>>\n", string(userJSON))
	fmt.Println("user1 saau khi Parse tu JSON =>>>\n", user1)

}
