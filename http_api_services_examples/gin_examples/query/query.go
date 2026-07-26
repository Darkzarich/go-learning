package query

import (
	"fmt"
	"net/http"
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

	// QueryMap parses key[subkey]=value pairs from the URL query string.
	g.GET("/posts", func(c *gin.Context) {
		filters := c.QueryMap("filters")

		fmt.Printf("filters: %v", filters)

		c.JSON(http.StatusOK, gin.H{
			"filters": filters,
		})
	})
}
