package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liperm/go-api/controllers"
)

func HandleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Hello).Methods("GET")
	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controllers.GetCustomerById).Methods("GET")
	log.Println("Ready to listen on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
