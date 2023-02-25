package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liperm/go-api/encryptions"
	"github.com/liperm/go-api/formatters"
	"github.com/liperm/go-api/models"
	"github.com/liperm/go-api/repository"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there!")
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers := repository.GetCustomers()
	json.NewEncoder(w).Encode(customers)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customer := repository.GetCustomerById(id)

	if customer.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	customer.Password = encryptions.EncriptData(customer.Password)

	id, message := repository.CreateCustomer(customer)

	if message != "OK" {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.FormatErrorResponse("Error creating customer: " + message)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
	succsessResponse := formatters.FormatCreationSuccessResponse(id)
	json.NewEncoder(w).Encode(succsessResponse)
}
