package repository

import (
	"github.com/liperm/go-api/constants"
	"github.com/liperm/go-api/database"
	"github.com/liperm/go-api/models"
)

func GetItems() []models.Item {
	var customers []models.Item
	database.DB.Find(&customers)
	return customers
}

func GetItemById(id int) models.Item {
	var item models.Item

	database.DB.Where("active = ?", true).First(&item, id)
	return item
}

func CreateItem(item models.Item) (int, string) {
	result := database.DB.Create(&item)

	if result.Error != nil {
		id := 0
		message := result.Error.Error()
		return id, message
	}

	id := item.ID
	message := constants.OK
	return id, message
}

func UpdateItem(item *models.Item) (int, string) {
	result := database.DB.Save(&item)
	if result.Error != nil {
		id := 0
		message := result.Error.Error()
		return id, message
	}
	return item.ID, constants.OK
}

func DeleteItem(item *models.Item) (int, string) {
	item.Active = false
	result := database.DB.Save(&item)

	if result.Error != nil {
		id := 0
		message := result.Error.Error()
		return id, message
	}
	return item.ID, constants.OK
}
