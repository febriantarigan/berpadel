package middleware

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// InitZap returns a configured zap logger and a Gin middleware
func InitZap() (*zap.Logger, gin.HandlerFunc) {
	logger, _ := zap.NewProduction()

	middleware := ginzap.Ginzap(logger, time.RFC3339, true)

	return logger, middleware
}

// StructuredLogger returns a middleware that track entry and exit of every request
func StructuredLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next() // Process the request

		latency := time.Since(start)

		logger.Info("incoming_request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.Duration("latency", latency),
			zap.String("ip", c.ClientIP()),
		)
	}
}

// ZapRecovery returns a middleware that recovers from panics and logs them via Zap
func ZapRecovery(logger *zap.Logger) gin.HandlerFunc {
	return ginzap.RecoveryWithZap(logger, true)
}
