package gomicroservices

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
	App.App.Run(":8081")
}
