package repository

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	acommon "go-microservices/a-common"

	"github.com/avinash-jh/go-microservice-common/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandlers struct {
	app *acommon.ApiHandlers
}

func NewProductHandlers(app *acommon.ApiHandlers) *ProductHandlers {
	return &ProductHandlers{
		app: app,
	}
}

func (h *ProductHandlers) AddProduct(c *gin.Context) {

	requestData, err := h.app.GetRequestData(c)

	if err != nil {
		acommon.SetErrorResponse(c, 500, err)
		return
	}

	var product models.Product

	if err := json.Unmarshal(requestData.Body, &product); err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid request body"))
		return
	}

	if err := h.app.DbConnection.Create(&product).Error; err != nil {
		log.Println("Failed to create product:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to create product"))
		return
	}

	acommon.SetSuccessResponse(c, 201,
		gin.H{
			"message": "Product created successfully",
			"data":    product,
		},
	)
}

func (h *ProductHandlers) GetProduct(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid product id"))
		return
	}

	var product models.Product

	if err := h.app.DbConnection.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			acommon.SetErrorResponse(c, 404, errors.New("product not found"))
			return
		}
		log.Println("Failed to get product:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to get product"))
		return
	}

	acommon.SetSuccessResponse(c, 200,
		gin.H{
			"data": product,
		},
	)
}

func (h *ProductHandlers) UpdateProduct(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid product id"))
		return
	}

	requestData, err := h.app.GetRequestData(c)

	if err != nil {
		acommon.SetErrorResponse(c, 500, err)
		return
	}

	var requestProduct models.Product

	if err := json.Unmarshal(requestData.Body, &requestProduct); err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid request body"))
		return
	}

	var product models.Product

	if err := h.app.DbConnection.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			acommon.SetErrorResponse(c, 404, errors.New("product not found"))
			return
		}
		log.Println("Failed to find product:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to find product"))
		return
	}

	product.Name = requestProduct.Name
	product.Description = requestProduct.Description
	product.Price = requestProduct.Price

	if err := h.app.DbConnection.Save(&product).Error; err != nil {
		log.Println("Failed to update product:", err)
		acommon.SetErrorResponse(c, 500, errors.New("failed to update product"))
		return
	}

	acommon.SetSuccessResponse(c, 200,
		gin.H{
			"message": "Product updated successfully",
			"data":    product,
		},
	)
}

func (h *ProductHandlers) DeleteProduct(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		acommon.SetErrorResponse(c, 400, errors.New("invalid product id"))
		return
	}

	result := h.app.DbConnection.Delete(&models.Product{}, id)

	if result.Error != nil {
		log.Println("Failed to delete product:", result.Error)
		acommon.SetErrorResponse(c, 500, errors.New("failed to delete product"))
		return
	}

	if result.RowsAffected == 0 {
		acommon.SetErrorResponse(c, 404, errors.New("product not found"))
		return
	}

	acommon.SetSuccessResponse(c, 200,
		gin.H{
			"message": "Product deleted successfully",
		},
	)
}
