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

func Create(w http.ResponseWriter, r *http.Request) {
	var people model.People
	logrus.Info("Insert people to DB")
	if err := json.NewDecoder(r.Body).Decode(&people); err != nil {
		logrus.Error("Decode error", err)
	}

	db.Connect().Create(&people)
	http.RedirectHandler("/", 200)
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	var peoples model.Peoples
	logrus.Info("Get all people")
	db.Connect().Find(&peoples)

	if err := json.NewEncoder(w).Encode(&peoples); err != nil {
		logrus.Error(err)
	}
}

func GetPeopleByID(w http.ResponseWriter, r *http.Request) {
	var people model.People
	vars := mux.Vars(r)
	id := vars["id"]
	logrus.Info("Find people by ID")
	db.Connect().First(&people, id)

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
		json.NewEncoder(w).Encode("Fail!!")
		return
	}

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.Error("Error: string convert", err)
		return
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
		return
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
		return
	}

	db.Connect().Delete(model.People{}, "id Like ?", uint(i))
}
