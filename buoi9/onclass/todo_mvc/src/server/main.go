package main

import (
	"net/http"
	"todo_mvc/routes"
)

func main() {
	http.ListenAndServe(":3000", routes.Routes())
}
