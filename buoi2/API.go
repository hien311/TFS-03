package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Calc(op string, a int64, b int64) int64 {
	switch op {
	case "mul":
		return a * b
	case "sub":
		return a - b
	case "add":
		return a + b
	case "div":
		if b == 0 {
			fmt.Println("Khong the chia cho 0")
			return 0
		} else {
			return a / b
		}
	default:
		return 0
	}
}

type Result struct {
	Result int64 `json:"result,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func calcResult(w http.ResponseWriter, r *http.Request) {
	result := Result{Result: Calc("mul", 3, 4)}
	resultJSON, _ := json.Marshal(result)
	fmt.Fprintln(w, string(resultJSON))
}

func user(w http.ResponseWriter, r *http.Request) {
	hien := User{Username: "Hien", Password: "12345678"}
	hienJSON, _ := json.MarshalIndent(hien, "", "   ")
	fmt.Println("http://localhost:8080/user")
	fmt.Fprintln(w, string(hienJSON))
	//json.NewEncoder(w).Encode(hien)
}

func handleRequests() {
	http.HandleFunc("/", calcResult)
	http.HandleFunc("/user", user)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequests()
}
