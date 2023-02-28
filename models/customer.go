package models

type Customer struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (customer *Customer) TableName() string {
	return "go_api.customer"
}
