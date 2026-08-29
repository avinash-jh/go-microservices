package repository

import (
	"fmt"
	"go-microservices/models"
	"log"

	"gorm.io/gorm"
)

func AddProduct(db *gorm.DB) {
	product := models.Product{
		Name:        "Laptop",
		Description: "Gaming Laptop",
		Price:       65000,
	}

	result := db.Create(&product)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Product created")
}

func GetProduct(db *gorm.DB) {
	product := models.Product{}
	result := db.First(&product, 1)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Product is", product)
}

func UpdateProduct(db *gorm.DB) {
	var product models.Product

	if err := db.First(&product, 1).Error; err != nil {
		log.Fatal("Product not found:", err)
	}

	result := db.Model(&product).Updates(models.Product{
		Name:  "Gaming Laptop Pro",
		Price: 75000,
	})

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Product updated successfully")
}

func DeleteProduct(db *gorm.DB) {
	result := db.Delete(&models.Product{}, 1)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		fmt.Println("No product found to delete")
		return
	}

	fmt.Println("Product deleted successfully")
}
