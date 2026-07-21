package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})

	r.POST("/users/:id/*action", func(c *gin.Context) {
		id := c.Param("id")
		action := c.Param("action")

		if action == "/error" {
			c.JSON(400, gin.H{
				"error": fmt.Sprintf("User ID: %s", id),
			})
		} else {
			c.JSON(200, gin.H{
				"id": id,
			})
		}
	})

	r.GET("/users/:id/string", func(c *gin.Context) {
		id := c.Param("id")

		c.String(200, "User ID: %s", id)
	})

	r.Run(":3000")
}
