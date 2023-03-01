package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/liperm/go-api/constants"
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
	id, _ := strconv.Atoi(vars["id"])
	customer := repository.GetCustomerById(id)

	if customer.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		notFoundResponse := formatters.FormatNotFoundResponse(id)
		json.NewEncoder(w).Encode(notFoundResponse)
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
	succsessResponse := formatters.FormatSuccessResponse(id)
	json.NewEncoder(w).Encode(succsessResponse)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	customer := repository.GetCustomerById(id)

	if customer.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		notFoundResponse := formatters.FormatNotFoundResponse(id)
		json.NewEncoder(w).Encode(notFoundResponse)
		return
	}
	json.NewDecoder(r.Body).Decode(&customer)
	id, message := repository.UpdateCustomer(&customer)

	if message != constants.OK {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.FormatErrorResponse("Error updating customer: " + message)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse := formatters.FormatSuccessResponse(id)
	json.NewEncoder(w).Encode(successResponse)
	w.WriteHeader(http.StatusOK)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	customer := repository.GetCustomerById(id)

	if customer.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		notFoundResponse := formatters.FormatNotFoundResponse(id)
		json.NewEncoder(w).Encode(notFoundResponse)
		return
	}
	id, message := repository.DeleteCustomer(&customer)

	if message != constants.OK {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.FormatErrorResponse("Error deleting customer: " + message)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse := formatters.FormatSuccessResponse(id)
	json.NewEncoder(w).Encode(successResponse)

	w.WriteHeader(http.StatusOK)
}
