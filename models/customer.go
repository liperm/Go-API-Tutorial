package models

type Customer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Passord string `json:"password"`
}
