package models

type Base struct {
	ID     int  `json:"id" gorm:"PrimaryKey"`
	Active bool `json:"active" gorm:"default:true"`
}
