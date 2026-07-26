package multipart_form

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/multipart_form")

	g.POST("/test_form", func(c *gin.Context) {
		message := c.PostForm("message")
		user := c.DefaultPostForm("user", "anonymous")

		c.JSON(200, gin.H{
			"status":  "success",
			"message": message,
			"user":    user,
		})
	})
}
