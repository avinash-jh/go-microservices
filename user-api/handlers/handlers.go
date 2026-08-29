package repository

import (
	"fmt"
	"go-microservices/models"
	"log"

	"gorm.io/gorm"
)

func AddUser(db *gorm.DB) {
	user := models.User{
		UserName:    "Avinash kumar",
		Email:       "avinashjha607@gmail.com",
		PhoneNumber: "7739876270",
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("user created")
}

func GetUser(db *gorm.DB) {
	user := models.User{}
	result := db.First(&user, 1)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("user is", user)
}

func UpdateUser(db *gorm.DB) {
	var user models.User

	if err := db.First(&user, 1).Error; err != nil {
		log.Fatal("user not found:", err)
	}

	result := db.Model(&user).Updates(models.User{
		Email:       "avinashsamrat607@gmail.com",
		PhoneNumber: "7723456480",
	})

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("user updated successfully")
}

func DeleteUser(db *gorm.DB) {
	result := db.Delete(&models.User{}, 1)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		fmt.Println("No user found to delete")
		return
	}

	fmt.Println("user deleted successfully")
}
