package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	uploadDir   = "./uploads"
	maxFileSize = 1024 * 1024 * 5 // 5MB
)

func RegisterRoutes(router *gin.Engine) {
	// Ensure upload directory exists
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		panic(err)
	}

	g := router.Group("/files")

	g.POST("/single/upload", func(c *gin.Context) {
		// Limit request size to 5MB
		if err := c.Request.ParseMultipartForm(maxFileSize); err != nil {
			c.JSON(400, gin.H{
				"error": "file size is too large",
			})
			return
		}

		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "cannot read file"})
			return
		}
		defer file.Close()

		// Limit file size to the same 5MB
		if header.Size > maxFileSize {
			c.JSON(400, gin.H{
				"error": "file size is too large",
			})
			return
		}

		// Validate file type by reading the first 512 bytes
		buffer := make([]byte, 512)
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			c.JSON(500, gin.H{"error": "cannot read file"})
			return
		}

		// Validate content type by reading the first 512 bytes or less
		contentType := http.DetectContentType(buffer[:bytesRead])
		if !isAllowedContentType(contentType) {
			c.JSON(400, gin.H{"error": "file type not allowed", "contentType": contentType})
			return
		}

		// Reset read pointer so can save later
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			c.JSON(500, gin.H{"error": "internal error"})
			return
		}

		// Validate file extension
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if !isAllowedExtension(ext) {
			c.JSON(400, gin.H{"error": "file extension not allowed"})
			return
		}

		// Generate unique name: timestamp + random suffix
		uniqueName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), uuid.New().String()[:8], ext)

		dst := filepath.Join(uploadDir, uniqueName)

		if err := c.SaveUploadedFile(header, dst); err != nil {
			c.JSON(500, gin.H{"error": "failed to save file"})
			return
		}

		pwd, err := os.Getwd()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		fullPath := filepath.Join(pwd, dst)

		c.JSON(201, gin.H{
			"filename": header.Filename,
			"size":     header.Size,
			"dst":      fullPath,
			"type":     contentType,
		})
	})
}

func isAllowedContentType(ct string) bool {
	allowed := map[string]bool{
		"image/jpeg":                true,
		"image/png":                 true,
		"image/gif":                 true,
		"application/pdf":           true,
		"text/plain":                true,
		"text/plain; charset=utf-8": true,
	}
	return allowed[ct]
}

func isAllowedExtension(ext string) bool {
	allowed := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".pdf":  true,
		".txt":  true,
	}
	return allowed[ext]
}
