package files

import (
	"errors"
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

const sniffBytes = 512

var allowedContentTypes = map[string]bool{
	"image/jpeg":                true,
	"image/png":                 true,
	"image/gif":                 true,
	"application/pdf":           true,
	"text/plain":                true,
	"text/plain; charset=utf-8": true,
}

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".pdf":  true,
	".txt":  true,
}

func RegisterRoutes(router *gin.Engine) error {
	// Ensure upload directory exists
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return fmt.Errorf("create upload dir: %w", err)
	}

	g := router.Group("/files")

	g.POST("/single/upload", func(c *gin.Context) {
		// Limit request body size to 5MB; oversized bodies fail at read time
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxFileSize))

		if err := c.Request.ParseMultipartForm(int64(maxFileSize)); err != nil {
			var maxBytesErr *http.MaxBytesError
			if errors.As(err, &maxBytesErr) {
				respondError(c, http.StatusRequestEntityTooLarge, "file size is too large")
				return
			}
			respondError(c, http.StatusBadRequest, "cannot parse multipart form: "+err.Error())
			return
		}
		// Clean up temp files written to disk by ParseMultipartForm
		defer c.Request.MultipartForm.RemoveAll()

		file, header, err := c.Request.FormFile("file")
		if err != nil {
			respondError(c, http.StatusBadRequest, "cannot read file")
			return
		}
		defer file.Close()

		// Sniff content type from the first 512 bytes
		buffer := make([]byte, sniffBytes)
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			respondError(c, http.StatusInternalServerError, "internal error")
			return
		}

		if bytesRead == 0 {
			respondError(c, http.StatusBadRequest, "file is empty")
			return
		}

		contentType := http.DetectContentType(buffer[:bytesRead])
		if !isAllowedContentType(contentType) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       "file type not allowed",
				"contentType": contentType,
			})
			return
		}

		// Validate file extension
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if !isAllowedExtension(ext) {
			respondError(c, http.StatusBadRequest, "file extension not allowed")
			return
		}

		// Generate unique name: timestamp + random suffix
		uniqueName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), uuid.New().String()[:8], ext)

		dst := filepath.Join(uploadDir, uniqueName)

		if err := c.SaveUploadedFile(header, dst); err != nil {
			respondError(c, http.StatusInternalServerError, "failed to save file")
			return
		}

		pwd, err := os.Getwd()
		if err != nil {
			respondError(c, http.StatusInternalServerError, err.Error())
			return
		}

		fullPath := filepath.Join(pwd, dst)

		c.JSON(http.StatusCreated, gin.H{
			"filename": header.Filename,
			"size":     header.Size,
			"dst":      fullPath,
			"type":     contentType,
		})
	})

	return nil
}

func respondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func isAllowedContentType(ct string) bool {
	return allowedContentTypes[ct]
}

func isAllowedExtension(ext string) bool {
	return allowedExtensions[ext]
}
