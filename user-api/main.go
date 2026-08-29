package main

import (
	"log"

	"go-microservices/database"
	repository "go-microservices/user-api/handlers"
)

func main() {

	db, err := database.ConnectAndMigrate()

	if err != nil {
		log.Fatal(err)
	}
	repository.AddUser(db)
}
