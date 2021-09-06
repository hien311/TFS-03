package routes

import (
	"github.com/gorilla/mux"

	controller "tfs-03db/controller"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.Path("/").HandlerFunc(controller.Home).Methods("GET")
	r.Path("/peoples").HandlerFunc(controller.GetAllPeople).Methods("GET")
	r.Path("/peoples/{id:[0-9]+}").HandlerFunc(controller.GetPeopleByID).Methods("GET")
	r.Path("/peoples").HandlerFunc(controller.Create).Methods("POST")
	r.Path("/peoples/{id:[0-9]+}").HandlerFunc((controller.UpdatePeople)).Methods("PUT")
	r.Path("/peoples/{id:[0-9]+}").HandlerFunc((controller.SoftDeletePeople)).Methods("POST")
	r.Path("/peoples/{id:[0-9]+}").HandlerFunc((controller.DeletePeople)).Methods("Delete")

	return r
}
