package acommon

import (
	"errors"

	"go-microservices/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiHandlers struct {
	GinApp       *gin.Engine
	DbConnection *gorm.DB
}

func NewApiHandler() *ApiHandlers {
	return &ApiHandlers{}
}

func (a *ApiHandlers) GetRequestData(c *gin.Context) (*RequestData, error) {

	value, exists := c.Get("requestData")

	if !exists {
		return nil, errors.New("request data not found")
	}

	requestData, ok := value.(*RequestData)

	if !ok {
		return nil, errors.New("invalid request data type")
	}

	return requestData, nil
}

func CreateApp() (*ApiHandlers, error) {

	// Connect to database
	db, err := database.ConnectAndMigrate()

	if err != nil {
		return nil, err
	}

	// Create Gin application
	app := gin.Default()

	// Create application handler
	apiHandler := NewApiHandler()

	// Register middleware
	app.Use(RequestMiddleware())
	app.Use(ResponseMiddleware())

	// Store dependencies
	apiHandler.GinApp = app
	apiHandler.DbConnection = db

	return apiHandler, nil
}
