package main

import (
	"log"

	"product-api/database"
	"product-api/repository"
)

func main() {

	db, err := database.ConnectAndMigrate()

	if err != nil {
		log.Fatal(err)
	}
	repository.AddProduct(db)
}
