package routes

import (
	acommon "go-microservices/a-common"
	repository "go-microservices/user-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *acommon.ApiHandlers) *gin.Engine {
	h := repository.NewUserHandlers(app)
	app.App.GET("/users/:id", h.GetUser)
	app.App.POST("/users", h.AddUser)
	app.App.PUT("/users/:id", h.UpdateUser)
	app.App.DELETE("/users/:id", h.DeleteUser)
	return app.App
}
