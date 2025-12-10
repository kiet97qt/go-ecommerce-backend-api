package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiKeyHeader = "X-API-Key"
const validAPIKey = "super-secret-key" // TODO: load from config

// AuthMiddleware kiểm tra API key trong header cho các request cần bảo vệ.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(apiKeyHeader)
		if key == "" || key != validAPIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid api key",
			})
			return
		}

		c.Next()
	}
}
