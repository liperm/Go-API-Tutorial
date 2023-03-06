package models

type CustomerItem struct {
	Base
	CustomerID int `json:"customer_id" gorm:"primaryKey"`
	ItemID     int `json:"item_id" gorm:"primaryKey"`
}

func (customer_item *CustomerItem) TableName() string {
	return "go_api.customer_item"
}
