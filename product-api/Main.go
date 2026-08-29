package main

import (
	"log"

	"go-microservices/database"
	repository "go-microservices/product-api/handlers"
)

func main() {

	db, err := database.ConnectAndMigrate()

	if err != nil {
		log.Fatal(err)
	}
	repository.DeleteProduct(db)
}
