package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/redirect")

	g.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/redirect/test2")
	})

	g.GET("/test2", func(c *gin.Context) {
		c.String(200, "This is GET /test2 endpoint")
	})

	g.POST("/test3", func(c *gin.Context) {
		// Redirects with 307 status code, which means that the request method is preserved
		c.Redirect(http.StatusTemporaryRedirect, "/redirect/test4")
	})

	g.POST("/test4", func(c *gin.Context) {
		c.String(200, "This is POST /test4 endpoint")
	})
}
