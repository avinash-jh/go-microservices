package acommon

import (
	"go-microservices/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiHandlers struct {
	App          *gin.Engine
	DbConnection *gorm.DB
}

func CreateApp() (*ApiHandlers, error) {
	db, err := database.ConnectAndMigrate()
	if err != nil {
		return nil, err
	}
	return &ApiHandlers{
		App:          gin.Default(),
		DbConnection: db,
	}, nil
}
