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

	//Customer
	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controllers.GetCustomerById).Methods("GET")
	router.HandleFunc("/customers/{id}/buying_list", controllers.GetCustomerBuyingList).Methods("GET")
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")

	//Item
	router.HandleFunc("/items", controllers.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", controllers.GetItemById).Methods("GET")
	router.HandleFunc("/items/code/{code}", controllers.GetItemByCode).Methods("GET")
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")

	log.Println("Ready to listen on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
