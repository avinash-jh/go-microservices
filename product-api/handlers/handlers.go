package repository

import (
	"fmt"
	acommon "go-microservices/a-common"
	"go-microservices/models"
	"log"

	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	app *acommon.ApiHandlers
}

func NewProductHandlers(app *acommon.ApiHandlers) *ProductHandlers {
	return &ProductHandlers{app: app}
}

func (h *ProductHandlers) AddProduct(c *gin.Context) {
	product := models.Product{
		Name:        "Laptop",
		Description: "Gaming Laptop",
		Price:       65000,
	}

	result := h.app.DbConnection.Create(&product)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Product created")
}

func (h *ProductHandlers) GetProduct(c *gin.Context) {
	product := models.Product{}
	result := h.app.DbConnection.First(&product, 1)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Product is", product)
}

func (h *ProductHandlers) UpdateProduct(c *gin.Context) {
	var product models.Product

	if err := h.app.DbConnection.First(&product, 1).Error; err != nil {
		log.Fatal("Product not found:", err)
	}

	result := h.app.DbConnection.Model(&product).Updates(models.Product{
		Name:  "Gaming Laptop Pro",
		Price: 75000,
	})

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Product updated successfully")
}

func (h *ProductHandlers) DeleteProduct(c *gin.Context) {
	result := h.app.DbConnection.Delete(&models.Product{}, 1)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		fmt.Println("No product found to delete")
		return
	}

	fmt.Println("Product deleted successfully")
}
