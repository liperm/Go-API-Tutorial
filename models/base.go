package models

type Base struct {
	ID     int  `json:"id"`
	Active bool `json:"active" gorm:"default:true"`
}
