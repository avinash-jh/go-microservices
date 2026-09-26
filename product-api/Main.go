package main

import (
	acommon "go-microservices/a-common"
	"go-microservices/product-api/routes"
)

func main() {
	App, err := acommon.CreateApp()
	if err != nil {
		panic(err)
	}
	routes.SetupRoutes(App)
	App.App.Run(":8080")
}
