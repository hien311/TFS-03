package routes

import (
	"github.com/gorilla/mux"

	controller "tfs-03db/controller"
	"tfs-03db/middleware"
)

func Routes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleware.ContentType)
	r.Use(middleware.AllowAccess)
	r.Path("/").HandlerFunc(controller.Home).Methods("GET")
	r.Path("/peoples").HandlerFunc(controller.GetAllPeople).Methods("GET")
	r.Path("/peoples/{id:[0-9]+}").HandlerFunc(controller.GetPeopleByID).Methods("GET")
	r.Path("/peoples").HandlerFunc(controller.Create).Methods("POST")
	r.Path("/peoples/update/{id:[0-9]+}").HandlerFunc((controller.UpdatePeople)).Methods("PUT", "OPTIONS")
	r.Path("/peoples/delete/{id:[0-9]+}").HandlerFunc((controller.SoftDeletePeople)).Methods("POST")
	r.Path("/peoples/delete/{id:[0-9]+}").HandlerFunc((controller.DeletePeople)).Methods("DELETE", "OPTIONS")

	return r
}
