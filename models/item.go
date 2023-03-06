package models

type Item struct {
	Base
	Name      string      `json:"name"`
	Code      string      `json:"code"`
	Quantity  int         `json:"quantity"`
	Customers []*Customer `json:"customers" gorm:"many2many:customer_item"`
	Value     float32     `json:"value"`
}

func (item *Item) TableName() string {
	return "go_api.item"
}
