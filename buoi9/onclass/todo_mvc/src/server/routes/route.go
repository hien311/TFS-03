package routes

import (
	"todo_mvc/controller"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(controller.CorAccept)
	r.Path("/tasks").HandlerFunc(controller.GetAllTasks).Methods("GET", "OPTIONS")
	r.Path("/tasks/add").HandlerFunc(controller.AddTask).Methods("POST")
	r.Path("/tasks/complete-all").HandlerFunc(controller.CompleteAllTask).Methods("POST")
	r.Path("/tasks/delete/{id:[0-9]+}").HandlerFunc(controller.DeleteTask).Methods("DELETE", "OPTIONS")
	r.Path("/tasks/delete/completed-task").HandlerFunc(controller.DeleteCompletedTask).Methods("DELETE", "OPTIONS")
	r.Path("/tasks/update/{id:[0-9]+}").HandlerFunc(controller.EditTask).Methods("PUT", "OPTIONS")
	return r
}
