package custom_error

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Status int    `json:"-"`
	Code   string `json:"code"`
	Msg    string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Msg
}

var (
	ErrNotFound = &AppError{
		Status: http.StatusNotFound,
		Code:   "not_found",
		Msg:    "Resource not found",
	}
	ErrBadRequest = &AppError{
		Status: http.StatusBadRequest,
		Code:   "bad_request",
		Msg:    "Invalid request",
	}
	ErrUnauthorized = &AppError{
		Status: http.StatusUnauthorized,
		Code:   "unauthorized",
		Msg:    "authentication required",
	}
)

func ErrorHandler(c *gin.Context) {
	// Need to use Next() here because this middleware is executed first, then handlers
	// so with this Next() it goes to the function below, then after error is saved with c.Error
	// it returns here and continues on
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	err := c.Errors.Last()
	var appErr *AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Status, gin.H{
			"success": false,
			"error":   appErr,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   gin.H{"code": "internal", "message": "Internal server error"},
		})
	}
}

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/custom-error")
	g.Use(ErrorHandler)

	g.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")

		if id == "0" {
			_ = c.Error(ErrNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id":   id,
				"name": "Some post",
			},
		})
	})
}
