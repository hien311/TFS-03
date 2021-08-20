package controller

import (
	"encoding/json"
	db "homework/database"
	"homework/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// cart manager
func GetAllCart(w http.ResponseWriter, r *http.Request) {
	logrus.Info("get all Order")
	var myCarts []models.Cart
	db.Connect().Where("Status = ?", "").Find(&myCarts)
	json.NewEncoder(w).Encode(myCarts)
}

func CreateCart(w http.ResponseWriter, r *http.Request) {
	logrus.Info("insert cart")
	var cart models.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		logrus.Error("Loi insert cart")
		return
	}
	db.Connect().Create(&cart)
}

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	logrus.Info("update cart")
	id := mux.Vars(r)["ID"]
	var newCart models.Cart
	err := json.NewDecoder(r.Body).Decode(&newCart)
	if err != nil {
		logrus.Error("Loi update")
	}
	newCart.ID = id
	db.Connect().Save(&newCart)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	logrus.Info("delete cart")
	id := mux.Vars(r)["ID"]
	var cart models.Cart
	db.Connect().First(&cart, id)
	cart.Status = "deleted"
	db.Connect().Save(&cart)
}
