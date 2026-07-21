package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
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

	r.Run(":3000")
}
