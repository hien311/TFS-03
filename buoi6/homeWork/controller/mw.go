package controller

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func CheckTokenMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := r.Cookie("logged")
		if r.URL.String() != "/" && token != nil {
			logrus.Info("you need to login")
			http.RedirectHandler("/", 200)
		} else {
			next.ServeHTTP(w, r)
		}
	})

}

func CheckContentTypeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != "application/json" && r.Method != "GET" {
			fmt.Fprint(w, "Wrong content type")
			return
		}
		next.ServeHTTP(w, r)
	})

}
