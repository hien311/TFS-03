package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func calcResult(w http.ResponseWriter, r *http.Request) {
	data := strings.Split(r.URL.String()[8:], "%20")
	symbol := string("")
	result, _ := strconv.ParseFloat(data[0], 64)
	for i, v := range data {
		if i > 0 {
			if v == "+" || v == "-" || v == "x" || v == "/" {
				symbol = v
			} else {
				temp, _ := strconv.ParseFloat(v, 64)
				switch symbol {
				case "+":
					result += temp
				case "-":
					result -= temp
				case "x":
					result *= temp
				case "/":
					result /= temp
				default:
					return
				}
			}
		}
	}
	resultJSON, _ := json.Marshal(result)
	fmt.Fprintln(w, string(resultJSON))
}

func UI(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "caculator.html", nil)
	//fmt.Fprintln(w, "calc")
	//json.NewEncoder(w).Encode(hien)
}

func handleRequests() {
	http.HandleFunc("/UI", UI)
	http.HandleFunc("/result/", calcResult)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func main() {
	fs := http.FileServer(http.Dir(""))
	http.Handle("/", http.StripPrefix(("/"), fs))
	handleRequests()
}
