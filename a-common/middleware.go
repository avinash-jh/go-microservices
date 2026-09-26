package acommon

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Body []byte
}

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			SetErrorResponse(c, 400, errors.New("failed to read request body"))
			return
		}

		c.Request.Body = io.NopCloser(
			bytes.NewBuffer(body),
		)

		requestData := &RequestData{
			Body: body,
		}

		c.Set("requestData", requestData)

		c.Next()
	}
}

type SuccessResponse struct {
	StatusCode int
	Data       any
}

func SetSuccessResponse(c *gin.Context, statusCode int, data any) {

	response := &SuccessResponse{
		StatusCode: statusCode,
		Data:       data,
	}

	c.Set("successResponse", response)
}

type ErrorResponse struct {
	StatusCode int
	ApiError   error
}

func SetErrorResponse(c *gin.Context, statusCode int, err error) {

	response := &ErrorResponse{
		StatusCode: statusCode,
		ApiError:   err,
	}

	c.Set("errorResponse", response)
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		// Check error response
		value, exists := c.Get("errorResponse")

		if exists {

			response, ok := value.(*ErrorResponse)

			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "invalid error response",
				})
				return
			}

			c.JSON(response.StatusCode, gin.H{
				"error": response.ApiError.Error(),
			})

			return
		}

		// Check success response
		value, exists = c.Get("successResponse")

		if !exists {
			return
		}

		response, ok := value.(*SuccessResponse)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "invalid success response",
			})
			return
		}

		c.JSON(response.StatusCode, response.Data)
	}
}
