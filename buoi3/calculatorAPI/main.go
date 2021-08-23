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
	result, _ := strconv.ParseFloat(data[0], 64)
	var result2 float64 = 0
	var sym2, sym1 string
	for i := range data {
		if data[i] == "+" || data[i] == "-" || data[i] == "x" || data[i] == "/" {
			sym1 = data[i]
			if data[i] == "+" || data[i] == "-" {
				switch sym2 {
				case "+":
					result += result2
				case "-":
					result -= result2
				}
				result2, _ = strconv.ParseFloat(data[i+1], 64)
				sym2 = data[i]
			}
		} else {
			value, _ := strconv.ParseFloat(data[i], 64)
			if result2 == 0 {
				result2 = value
			} else {
				switch sym1 {
				case "x":
					result2 *= value
				case "/":
					result2 /= value
				}
			}
		}
	}
	switch sym2 {
	case "+":
		result += result2
	case "-":
		result -= result2
	case "*":
		result *= result2
	case "/":
		result /= result2
	default:
		result = result2
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
