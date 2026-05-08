package httputil

import (
	"fmt"
	"log"
	"net/http"
	"users-service/pkg/apperror"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Translates a service error into an HTTP status and message.
func RespondError(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	msg := "internal server error"

	if appError, ok := err.(*apperror.AppError); ok {
		switch appError.Kind {
		case apperror.KindNotFound:
			status = http.StatusNotFound
		case apperror.KindInvalidInput:
			status = http.StatusBadRequest
		case apperror.KindAlreadyExists:
			status = http.StatusConflict
		}
		msg = appError.Message
	}

	c.JSON(status, gin.H{"error": msg})
	log.Printf("HTTP error: %v\n", err)
}

func RespondOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// Builds human-readable error message from validation errors
func RespondValidationError(c *gin.Context, err error) {
	var msg string

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			msg = fmt.Sprintf("field '%s' failed validation: %s", fe.Field(), fe.Tag())
			break
		}
	} else {
		msg = "invalid JSON body"
	}

	RespondError(c, apperror.NewInvalidInput(msg))
}
