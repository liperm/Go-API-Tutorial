package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there!")
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there from getCustomer!")
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Hello there from getCustomerById!", vars["id"])
}
