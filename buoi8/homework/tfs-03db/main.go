package main

import (
	"net/http"
	"tfs-03db/db"
	routes "tfs-03db/routes"
)

func main() {
	database := db.Connect()
	database.ServeInit()
	http.ListenAndServe(":3000", routes.Routes())
}
