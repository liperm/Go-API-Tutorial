package main

import (
	"log"

	"github.com/liperm/go-api/database"
	"github.com/liperm/go-api/routes"
)

func main() {
	log.Println("Initializing...")
	database.Connect()
	routes.HandleRequests()
}
