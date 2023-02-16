package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	conncetion := "host=localhost user=root dbname=root password=root port=5438 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conncetion))
	if err != nil {
		log.Panic("Database connection error")
	}
}
