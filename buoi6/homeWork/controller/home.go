package controller

import (
	"encoding/json"
	db "homework/database"
	"homework/models"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Login page")
	json.NewEncoder(w).Encode("this is homepage")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var temp models.User
	json.NewDecoder(r.Body).Decode(&user)

	db.Connect().First(&temp, user.UserName)
	if temp.UserName == "" {
		logrus.Error("wrong username")
		return
	}
	if temp.Password == user.Password {
		cookie := http.Cookie{
			Name:    "logged",
			Value:   "true",
			Path:    "/",
			Expires: time.Now().AddDate(0, 0, 10),
		}
		http.SetCookie(w, &cookie)
	}
}
