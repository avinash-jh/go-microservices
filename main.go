package main

import (
	acommon "go-microservices/a-common"
	productroutes "go-microservices/product-api/routes"
	userroutes "go-microservices/user-api/routes"
)

func main() {

	App, err := acommon.CreateApp()
	if err != nil {
		panic(err)
	}
	productroutes.SetupRoutes(App)
	userroutes.SetupRoutes(App)
	App.GinApp.Run(":8081")
}
