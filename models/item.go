package models

type Item struct {
	Base
	Name     string `json:"name"`
	Code     int    `json:"code"`
	Quantity int    `json:"quantity"`
}

func (item *Item) TableName() string {
	return "go_api.item"
}
