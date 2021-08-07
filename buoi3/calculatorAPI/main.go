package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

type Result struct {
	Result int64 `json:"result,omitempty"`
}

func takeReq(w http.ResponseWriter, r *http.Request) {
	string := r.Method
	fmt.Fprintln(w, string)
}

func calcResult(w http.ResponseWriter, r *http.Request) {
	hehe := r.Body
	resultJSON, _ := json.Marshal(hehe)
	fmt.Fprintln(w, string(resultJSON))
}

func UI(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "caculator.html", nil)
	//fmt.Fprintln(w, "calc")
	//json.NewEncoder(w).Encode(hien)
}

func handleRequests() {
	http.HandleFunc("/UI", UI)
	http.HandleFunc("/sendReq", takeReq)
	http.HandleFunc("/result", calcResult)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func main() {
	fs := http.FileServer(http.Dir(""))
	http.Handle("/", http.StripPrefix(("/"), fs))
	handleRequests()
}
