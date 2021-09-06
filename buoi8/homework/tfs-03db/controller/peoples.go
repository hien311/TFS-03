package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	db "tfs-03db/db"
	"tfs-03db/model"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func badRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var people model.People
	logrus.Info("Insert people to DB")
	if err := json.NewDecoder(r.Body).Decode(&people); err != nil {
		logrus.Error("Decode error", err)
	}

	if err := db.Connect().Create(&people); err != nil {
		logrus.Error("Error when insert to DB", err)
		badRequest(w, r)
	}

	http.RedirectHandler("/", 200)
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	var peoples model.Peoples
	logrus.Info("Get all people")
	if err := db.Connect().Find(&peoples); err != nil {
		logrus.Error(err)
	}
	if err := json.NewEncoder(w).Encode(&peoples); err != nil {
		logrus.Error(err)
	}
}

func GetPeopleByID(w http.ResponseWriter, r *http.Request) {
	var people model.People
	vars := mux.Vars(r)
	id := vars["id"]
	logrus.Info("Find people by ID")
	if err := db.Connect().First(&people, id); err != nil {
		logrus.Error(err)
	}
	if err := json.NewEncoder(w).Encode(people); err != nil {
		logrus.Error(err)
	}
}

func UpdatePeople(w http.ResponseWriter, r *http.Request) {
	var people model.People
	id := mux.Vars(r)["id"]
	logrus.Info("Update people")
	if err := json.NewDecoder(r.Body).Decode(&people); err != nil {
		logrus.Error("Update fail")
		badRequest(w, r)
	}

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Error("Error: string convert", err)
	}

	people.ID = uint(i)
	db.Connect().Save(&people)
}

func SoftDeletePeople(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	logrus.Info("Delete people")
	var people model.People

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Error("Error: string convert", err)
	}

	if err := db.Connect().First(&people, uint(i)); err != nil {
		logrus.Error("Error: find", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	people.Deleted = true
	db.Connect().Save(&people)
}

func DeletePeople(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	logrus.Info("Delete people")

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Error("Error: string convert", err)
	}

	db.Connect().Delete(model.People{}, "id Like ?", uint(i))
}
