package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"api-gateway/logger"  
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		start := time.Now()
		
		c.Next()

		duration := time.Since(start)
		logger.Info("Incoming request",
			map[string]interface{}{
				"method":   c.Request.Method,
				"path":     c.Request.URL.Path,
				"status":   c.Writer.Status(),
				"duration": duration,
			})
	}
}
