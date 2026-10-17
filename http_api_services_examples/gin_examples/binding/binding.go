package binding

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/binding")

	type User struct {
		Name     string
		LastName string
		Age      int
	}

	users := make([]User, 0)

	g.POST("/simple/users", func(c *gin.Context) {
		var userDTO struct {
			Name     string `json:"name" binding:"required,min=2,max=20"`
			LastName string `json:"last_name" binding:"required,min=2,max=20"`
			Age      int    `json:"age" binding:"required,min=18"`
		}

		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})

			return
		}

		for _, user := range users {
			if user.Name == userDTO.Name && user.LastName == userDTO.LastName && user.Age == userDTO.Age {
				c.JSON(409, gin.H{
					"error": "user already exists",
				})

				return
			}
		}

		users = append(users, User{
			Name:     userDTO.Name,
			LastName: userDTO.LastName,
			Age:      userDTO.Age,
		})

		c.JSON(200, gin.H{
			"name":      userDTO.Name,
			"last_name": userDTO.LastName,
			"age":       userDTO.Age,
		})
	})

	g.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
}
