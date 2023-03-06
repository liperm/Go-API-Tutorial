package repository

import (
	"github.com/liperm/go-api/constants"
	"github.com/liperm/go-api/database"
	"github.com/liperm/go-api/models"
)

func GetCustomers() []models.Customer {
	var customers []models.Customer
	database.DB.Model(&customers).Preload("Items", "active = ?", true).Find(&customers)
	return customers
}

func GetCustomerById(id int) models.Customer {
	var customer models.Customer

	database.DB.Model(&customer).Preload("Items", "active = ?", true).First(&customer, id).Where("active = ?", true)
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

func UpdateCustomer(customer *models.Customer) (int, string) {
	result := database.DB.Save(&customer)
	if result.Error != nil {
		id := 0
		message := result.Error.Error()
		return id, message
	}
	return customer.ID, constants.OK
}

func DeleteCustomer(customer *models.Customer) (int, string) {
	customer.Active = false
	result := database.DB.Save(&customer)

	if result.Error != nil {
		id := 0
		message := result.Error.Error()
		return id, message
	}
	return customer.ID, constants.OK
}
