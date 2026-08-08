package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const token_header = "Token"
const fake_secret = "123456"

func checkToken(c *gin.Context) {
	header := c.GetHeader(token_header)

	if header != fake_secret {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	c.Next()

}

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/middleware", checkToken)

	g.GET("/test", func(c *gin.Context) {
		c.String(200, "This is /test endpoint, passed fake auth")
	})

	g.GET("/test2", func(c *gin.Context) {
		c.String(200, "This is /test2 endpoint, passed fake auth")
	})
}
