package main

import (
	"net/http"
	routes "tfs-03db/routes"
)

func Init() {
	http.ListenAndServe(":3000", routes.Routes())
}

func main() {
}
