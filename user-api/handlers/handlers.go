package repository

import (
	"fmt"
	acommon "go-microservices/a-common"
	"go-microservices/models"
	"log"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	app *acommon.ApiHandlers
}

func NewUserHandlers(app *acommon.ApiHandlers) *UserHandlers {
	return &UserHandlers{app: app}
}

func (h *UserHandlers) AddUser(c *gin.Context) {
	user := models.User{
		UserName:    "Avinash kumar",
		Email:       "avinashjha607@gmail.com",
		PhoneNumber: "7739876270",
	}
	result := h.app.DbConnection.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("user created")
}

func (h *UserHandlers) GetUser(c *gin.Context) {
	user := models.User{}
	result := h.app.DbConnection.First(&user, 1)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("user is", user)
}

func (h *UserHandlers) UpdateUser(c *gin.Context) {
	var user models.User

	if err := h.app.DbConnection.First(&user, 1).Error; err != nil {
		log.Fatal("user not found:", err)
	}

	result := h.app.DbConnection.Model(&user).Updates(models.User{
		Email:       "avinashsamrat607@gmail.com",
		PhoneNumber: "7723456480",
	})

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("user updated successfully")
}

func (h *UserHandlers) DeleteUser(c *gin.Context) {
	result := h.app.DbConnection.Delete(&models.User{}, 1)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		fmt.Println("No user found to delete")
		return
	}

	fmt.Println("user deleted successfully")
}
