package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func ContentType(next http.Handler) http.Handler {
	jsonContentType := "application/json"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != jsonContentType && r.Method == "POST" {
			logrus.Error("Wrong content-type")
			json.NewEncoder(w).Encode("Require application/json Content-Type")
		}
		next.ServeHTTP(w, r)
	})
}

func AllowAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		next.ServeHTTP(w, r)
	})
}
