package main

import (
	"log"

	"go-microservices/database"
	"go-microservices/product-api/repository"
)

func main() {

	db, err := database.ConnectAndMigrate()

	if err != nil {
		log.Fatal(err)
	}
	repository.DeleteProduct(db)
}
