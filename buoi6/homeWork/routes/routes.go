package routes

import (
	controller "homework/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Connect() {
	route := mux.NewRouter().StrictSlash(true)
	route.Use(controller.CheckContentTypeHeader)
	route.Use(controller.CheckTokenMW)
	//
	route.Path("/").HandlerFunc(controller.Homepage).Methods("GET")
	route.Path("/").HandlerFunc(controller.Login).Methods("POST")
	// route myproduce
	route.Path("/my-produces").HandlerFunc(controller.MyProduces).Methods("GET")
	route.Path("/my-produces").HandlerFunc(controller.CreateProduce).Methods("POST")
	route.Path("/my-produces/{ID:[0-9]+}").HandlerFunc(controller.UpdateProduce).Methods("PATCH")
	route.Path("/my-produces/{ID:[0-9]-}").HandlerFunc(controller.DeleteProduce).Methods("DELETE")
	//routes Cart
	route.Path("/cart").HandlerFunc(controller.GetAllCart).Methods("GET")
	route.Path("/cart").HandlerFunc(controller.CreateCart).Methods("POST")
	route.Path("/cart/{ID:[0-9]+}").HandlerFunc(controller.UpdateCart).Methods("PATCH")
	route.Path("/cart/{ID:[0-9]-}").HandlerFunc(controller.DeleteCart).Methods("DELETE")
	// Pay
	route.Path("/pay").HandlerFunc(controller.PayPage).Methods("GET")
	http.ListenAndServe(":3000", route)
}
