package main

import (
	"log"

	"github.com/liperm/go-api/routes"
)

func main() {
	log.Println("Initializing...")
	routes.HandleRequests()
}
