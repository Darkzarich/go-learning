package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/test_form", func(c *gin.Context) {
		message := c.PostForm("message")
		user := c.DefaultPostForm("user", "anonymous")

		c.JSON(200, gin.H{
			"status":  "success",
			"message": message,
			"user":    user,
		})
	})

	// To test: curl -X POST -d "message=something&user=me" "localhost:3000/test_form"
	r.Run(":3000")
}
