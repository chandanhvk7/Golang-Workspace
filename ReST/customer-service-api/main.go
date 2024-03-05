package main

import (
	"api/controllers"
	"api/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middlewares.LogRequestMiddleware)
	r.Use(middlewares.RejectNonJsonRequest)
	r.HandleFunc("/", controllers.Home)
	api:=r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.RejectNonJsonRequest)
	api.Use(middlewares.JsonResponseMiddleware)
	api.HandleFunc("/customers", controllers.HandleGetAllCustomers).Methods("GET")
	api.HandleFunc("/customers/{id}", controllers.HandleGetOneCustomers).Methods("GET")
	api.HandleFunc("/customers",controllers.HandlePostOneCustomer).Methods("POST")
	api.HandleFunc("/customers/{id}",controllers.HandlePutCustomer).Methods("PUT")
	fmt.Println("Server running in port 7777")
	http.ListenAndServe(":7777", r)
}
