package routes

import (
	acommon "go-microservices/a-common"
	repository "go-microservices/product-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *acommon.ApiHandlers) *gin.Engine {
	h := repository.NewProductHandlers(app)
	app.App.GET("/products/:id", h.GetProduct)
	app.App.POST("/products", h.AddProduct)
	app.App.PUT("/products/:id", h.UpdateProduct)
	app.App.DELETE("/products/:id", h.DeleteProduct)
	return app.App
}
