package controller

import (
	"encoding/json"
	db "homework/database"
	"homework/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// route quan ly san pham
func MyProduces(w http.ResponseWriter, r *http.Request) {
	logrus.Info("get all Produces")
	var myProduces models.Produces
	db.Connect().Find(&myProduces)
	json.NewEncoder(w).Encode(myProduces)
}

func CreateProduce(w http.ResponseWriter, r *http.Request) {
	logrus.Info("insert produce")
	var produce models.Produce
	err := json.NewDecoder(r.Body).Decode(&produce)
	if err != nil {
		logrus.Error("Loi insert produce")
		return
	}
	db.Connect().Create(&produce)
}

func UpdateProduce(w http.ResponseWriter, r *http.Request) {
	logrus.Info("update Produce")
	id := mux.Vars(r)["ID"]
	var newProduce models.Produce
	err := json.NewDecoder(r.Body).Decode(&newProduce)
	if err != nil {
		logrus.Error("Loi update")
	}
	newProduce.ID = id
	db.Connect().Save(&newProduce)

}

func DeleteProduce(w http.ResponseWriter, r *http.Request) {
	logrus.Info("delete produce")
	id := mux.Vars(r)["ID"]
	var produce models.Produce
	db.Connect().First(&produce, id)
	db.Connect().Delete(&produce)
}
