package repository

import (
	"github.com/liperm/go-api/constants"
	"github.com/liperm/go-api/database"
	"github.com/liperm/go-api/models"
)

func GetCustomers() []models.Customer {
	var customers []models.Customer
	database.DB.Find(&customers)
	return customers
}

func GetCustomerById(id string) models.Customer {
	var customer models.Customer

	database.DB.First(&customer, id)
	return customer
}

func CreateCustomer(customer models.Customer) (int, string) {
	result := database.DB.Create(&customer)

	if result.Error != nil {
		id := 0
		message := result.Error.Error()
		return id, message
	}

	id := customer.ID
	message := constants.OK
	return id, message
}
