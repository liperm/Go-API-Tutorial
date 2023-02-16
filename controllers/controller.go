package controllers

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there!")
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there!")
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
}
