package main

import (
	db "homework/database"
	"homework/routes"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info(".....Starting.......")
	db.CreateTable()
	routes.Connect()
}
