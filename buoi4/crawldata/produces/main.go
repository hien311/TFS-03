package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	u, _ := http.Get("https://template-homedecor.onshopbase.com/collections/new-arrivals")
	res, _ := ioutil.ReadAll(u.Body)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(res))
	})
	http.ListenAndServe(":3000", nil)
	//data := string(res)

	//data2 := make(map[string]interface{})
	//json.Unmarshal(res, &data2)
	//fmt.Println(data2["code"])
}
