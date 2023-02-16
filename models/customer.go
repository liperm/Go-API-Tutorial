package models

type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Passord string `json:"password"`
}

func (customer *Customer) TableName() string {
	return "go_api.customer"
}
