package routes

import (
	acommon "go-microservices/a-common"
	repository "go-microservices/user-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *acommon.ApiHandlers) *gin.Engine {
	h := repository.NewUserHandlers(app)
	app.GinApp.GET("/users/:id", h.GetUser)
	app.GinApp.POST("/users", h.AddUser)
	app.GinApp.PUT("/users/:id", h.UpdateUser)
	app.GinApp.DELETE("/users/:id", h.DeleteUser)
	return app.GinApp
}
