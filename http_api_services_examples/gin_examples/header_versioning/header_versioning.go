package header_versioning

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VersionMiddleware(c *gin.Context) {
	version := c.Request.Header.Get("Accept-Version")
	if version == "" {
		version = "v1"
	}

	c.Set("version", version)
	c.Next()
}

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/header-versioning")
	g.Use(VersionMiddleware)

	g.GET("/hello", func(c *gin.Context) {
		version := c.GetString("version")

		switch version {
		case "v2":
			c.JSON(http.StatusOK, gin.H{
				"version": "v2",
				"data":    []gin.H{},
				"meta":    gin.H{},
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"version": "v1",
				"data":    []gin.H{},
			})
		}
	})
}
