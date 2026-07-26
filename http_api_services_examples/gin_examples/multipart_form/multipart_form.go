package multipart_form

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/multipart_form")

	g.POST("/test_form", func(c *gin.Context) {
		message := c.PostForm("message")
		user := c.DefaultPostForm("user", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": message,
			"user":    user,
		})
	})

	// PostFormMap — parses key[subkey]=value pairs from the request body.
	g.POST("/posts", func(c *gin.Context) {
		filters := c.PostFormMap("filters")

		fmt.Printf("filters: %v", filters)

		c.JSON(http.StatusOK, gin.H{
			"filters": filters,
		})
	})
}
