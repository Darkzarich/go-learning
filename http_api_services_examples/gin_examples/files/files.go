package files

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/files")

	g.POST("/single/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		dst := filepath.Join("./", filepath.Base(file.Filename))

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		pwd, err := os.Getwd()
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		fullPath := filepath.Join(pwd, dst)

		c.JSON(200, gin.H{
			"filename": file.Filename,
			"size":     file.Size,
			"dst":      fullPath,
			"type":     file.Header.Get("Content-Type"),
		})
	})
}
