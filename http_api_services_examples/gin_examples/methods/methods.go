package methods

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/methods")

	g.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})

	g.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})

	g.PUT("/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})
	})

	g.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete",
		})
	})

	g.PATCH("/patch", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "patch",
		})
	})

	g.HEAD("/head", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "head",
		})
	})

	g.OPTIONS("/options", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "options",
		})
	})
}
