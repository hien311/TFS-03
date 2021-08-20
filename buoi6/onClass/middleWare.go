package main

import (
	"fmt"
	"mw/login"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID   string
	Name string
	Age  string
}

var studentList []Student

func main() {
	router := mux.NewRouter().StrictSlash(true)
	privateRouter := router.NewRoute().Subrouter()
	privateRouter.Use(login.CheckTokenMW)
	upRouter := router.NewRoute().Subrouter()
	upRouter.Use(contentTypeCheck)
	// CRUD
	router.Path("/").HandlerFunc(Login).Methods("GET")
	privateRouter.Path("/students").HandlerFunc(Get).Methods("GET")
	upRouter.Path("/students").HandlerFunc(Create).Methods("POST")
	upRouter.Path("/students/{ID:[0-9]+}").HandlerFunc(Update).Methods("PATCH")
	upRouter.Path("/students/{ID:[0-9]+}").HandlerFunc(Delete).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}

func contentTypeCheck(next http.Handler) http.Handler {
	jsonContentType := "application/json"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HeaderContent := r.Header.Get("Content-Type")
		if HeaderContent != jsonContentType {
			fmt.Fprintf(w, "wrong")
		}
		next.ServeHTTP(w, r)
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("token", "123")
	fmt.Fprint(w, "Login here")
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, studentList)
}

func Create(w http.ResponseWriter, r *http.Request) {
	student := Student{Name: r.FormValue("Name"), Age: r.FormValue("Age")}
	w.Header().Set("Content-Type", "application/json")
	student.ID = strconv.Itoa(len(studentList))
	studentList = append(studentList, student)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["ID"])
	studentList = append(studentList[:id], studentList[id+1:]...)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["ID"])
	studentList[id].ID = vars["ID"]
	keys := []string{"Name", "Age"}

	for _, v := range keys {
		if r.FormValue(v) != "" {
			switch v {
			case "Name":
				studentList[id].Name = r.FormValue(v)
			case "Age":
				studentList[id].Age = r.FormValue(v)
			}
		}
	}
}
