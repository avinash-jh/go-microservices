package routes

import (
	acommon "go-microservices/a-common"
	repository "go-microservices/product-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *acommon.ApiHandlers) *gin.Engine {

	productHandlers := repository.NewProductHandlers(app)

	app.GinApp.GET("/products/:id", productHandlers.GetProduct)
	app.GinApp.POST("/products", productHandlers.AddProduct)
	app.GinApp.PUT("/products/:id", productHandlers.UpdateProduct)
	app.GinApp.DELETE("/products/:id", productHandlers.DeleteProduct)

	return app.GinApp
}
