package main

import (
	acommon "go-microservices/a-common"
	"go-microservices/product-api/routes"
	"log"
)

func main() {

	app, err := acommon.CreateApp()
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupRoutes(app)

	if err := app.GinApp.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
