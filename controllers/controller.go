package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/liperm/go-api/database"
	"github.com/liperm/go-api/models"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there!")
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer
	database.DB.Find(&customers)
	json.NewEncoder(w).Encode(customers)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
}
