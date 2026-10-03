package repository

import (
	"encoding/json"
	"errors"
	acommon "go-microservices/a-common"
	"log"

	"github.com/avinash-jh/go-microservice-common/models"
	"github.com/avinash-jh/go-microservice-common/security"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandlers struct {
	app *acommon.ApiHandlers
}

func NewUserHandlers(app *acommon.ApiHandlers) *UserHandlers {
	return &UserHandlers{
		app: app,
	}
}

func (h *UserHandlers) AddUser(c *gin.Context) {

	requestData, err := h.app.GetRequestData(c)
	if err != nil {
		acommon.SetErrorResponse(c, 500, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(requestData.Body, &user); err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid request body"))
		return
	}

	hashedPassword, err := security.HashPassword(user.Password)

	if err != nil {
		log.Println("Failed to hash password:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to process password"))
		return
	}

	user.Password = hashedPassword

	if err := h.app.DbConnection.Create(&user).Error; err != nil {
		log.Println("Failed to create user:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to create user"))
		return
	}

	acommon.SetSuccessResponse(c, 201,
		gin.H{
			"message": "User created successfully",
			"data":    user,
		},
	)
}

func (h *UserHandlers) GetUser(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		acommon.SetErrorResponse(c, 400, errors.New("user id is blank"))
		return
	}

	var user models.User

	if err := h.app.DbConnection.First(&user, "id = ?", id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			acommon.SetErrorResponse(c, 404, errors.New("user not found"))
			return
		}

		log.Println("Failed to get user:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to get user"))
		return
	}

	acommon.SetSuccessResponse(c, 200,
		gin.H{
			"data": user,
		},
	)
}

func (h *UserHandlers) UpdateUser(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		acommon.SetErrorResponse(c, 400, errors.New("user id is blank"))
		return
	}

	requestData, err := h.app.GetRequestData(c)
	if err != nil {
		acommon.SetErrorResponse(c, 500, err)
		return
	}

	var updates map[string]interface{}

	if err := json.Unmarshal(requestData.Body, &updates); err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid request body"))
		return
	}

	if len(updates) == 0 {
		acommon.SetErrorResponse(c, 400, errors.New("no fields to update"))
		return
	}

	if _, exists := updates["id"]; exists {
		acommon.SetErrorResponse(c, 400, errors.New("id cannot be updated"))
		return
	}

	if _, exists := updates["password"]; exists {
		acommon.SetErrorResponse(c, 400, errors.New("password cannot be updated"))
		return
	}

	result := h.app.DbConnection.
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		log.Println("Failed to update user:", result.Error)

		acommon.SetErrorResponse(
			c,
			500,
			errors.New("failed to update user"),
		)
		return
	}

	var user models.User

	if err := h.app.DbConnection.
		First(&user, "id = ?", id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			acommon.SetErrorResponse(
				c,
				404,
				errors.New("user not found"),
			)
			return
		}

		log.Println("Failed to get updated user:", err)

		acommon.SetErrorResponse(
			c,
			500,
			errors.New("failed to get updated user"),
		)
		return
	}

	acommon.SetSuccessResponse(c, 200, gin.H{
		"message": "User updated successfully",
		"data":    user,
	})
}

func (h *UserHandlers) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		acommon.SetErrorResponse(c, 400, errors.New("user id is blank"))
		return
	}

	result := h.app.DbConnection.Delete(&models.User{}, id)

	if result.Error != nil {
		log.Println("Failed to delete user:", result.Error)
		acommon.SetErrorResponse(c, 500, errors.New("failed to delete user"))
		return
	}

	if result.RowsAffected == 0 {
		acommon.SetErrorResponse(c, 404, errors.New("user not found"))
		return
	}

	acommon.SetSuccessResponse(c, 200,
		gin.H{
			"message": "User deleted successfully",
		},
	)
}
