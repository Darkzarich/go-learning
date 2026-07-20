package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})

	r.PUT("/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete",
		})
	})

	r.PATCH("/patch", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "patch",
		})
	})

	r.HEAD("/head", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "head",
		})
	})

	r.OPTIONS("/options", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "options",
		})
	})

	r.Run(":3000")
}
