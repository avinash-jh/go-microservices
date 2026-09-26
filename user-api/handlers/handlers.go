package repository

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	acommon "go-microservices/a-common"
	"go-microservices/models"

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

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid user id"))
		return
	}

	var user models.User

	if err := h.app.DbConnection.First(&user, id).Error; err != nil {
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

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid user id"))
		return
	}

	requestData, err := h.app.GetRequestData(c)

	if err != nil {
		acommon.SetErrorResponse(c, 500, err)
		return
	}

	var requestUser models.User

	if err := json.Unmarshal(requestData.Body, &requestUser); err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid request body"))
		return
	}

	var user models.User

	if err := h.app.DbConnection.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			acommon.SetErrorResponse(c, 404, errors.New("user not found"))
			return
		}
		log.Println("Failed to find user:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to find user"))
		return
	}

	user.UserName = requestUser.UserName
	user.Email = requestUser.Email
	user.PhoneNumber = requestUser.PhoneNumber

	if err := h.app.DbConnection.Save(&user).Error; err != nil {
		log.Println("Failed to update user:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to update user"))
		return
	}

	acommon.SetSuccessResponse(c, 200,
		gin.H{
			"message": "User updated successfully",
			"data":    user,
		},
	)
}

func (h *UserHandlers) DeleteUser(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid user id"))
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
