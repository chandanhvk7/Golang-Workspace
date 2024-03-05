package controllers

import (
	"api/dao"
	"api/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Customer Service end Point Here")
}

func HandleGetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := dao.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

func HandleGetOneCustomers(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	c := dao.GetOneCustomer(id)
	if c == nil {
		w.WriteHeader(http.StatusNotFound)
		err := ErrorMessage{fmt.Sprintf("ID:%d not found!", id)}
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(*c)
	}
}

func HandlePostOneCustomer(w http.ResponseWriter, r *http.Request) {
	var cust model.Customer
	json.NewDecoder(r.Body).Decode(&cust)
	cust.Id = dao.AddCustomer(cust)
	w.WriteHeader(http.StatusCreated) //201
	json.NewEncoder(w).Encode(cust)
}

func HandlePutCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var cust model.Customer
	json.NewDecoder(r.Body).Decode(&cust)
	c:= dao.UpdateCustomer(id, cust)
	if c == nil {
		w.WriteHeader(http.StatusNotFound)
		err := ErrorMessage{fmt.Sprintf("ID:%d not found!", id)}
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(*c)
	}
}
