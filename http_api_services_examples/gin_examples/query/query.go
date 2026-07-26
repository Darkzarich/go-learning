package query

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/query")

	g.GET("/users", func(c *gin.Context) {
		page := c.Query("page")
		limit := c.Query("limit")
		sort := c.DefaultQuery("sort", "asc")

		if sort != "asc" && sort != "desc" {
			sort = "asc"
		}

		intPage, _ := strconv.Atoi(page)
		intLimit, _ := strconv.Atoi(limit)

		c.JSON(200, gin.H{
			"page":  intPage,
			"limit": intLimit,
			"sort":  sort,
		})
	})
}
