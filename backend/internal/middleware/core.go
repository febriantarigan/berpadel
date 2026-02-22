package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// cache request body from gin request, so it can be reused by other process
func CacheRequestBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			// Restore the body for the next person
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			c.Set("cachedBody", string(bodyBytes))
		}
		c.Next()
	}
}

// middleware to handle request with no route (e.g wrong endpoint)
func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    "PAGE_NOT_FOUND",
		"message": "The requested endpoint does not exist",
	})
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process everything first

		// If any errors were collected during the request
		if len(c.Errors) > 0 {
			c.JSON(c.Writer.Status(), gin.H{
				"errors":  c.Errors.Errors(),
				"success": false,
			})
		}
	}
}
