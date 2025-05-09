package middlewares

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("NX_S3_TOKEN")
	formattedSecret := fmt.Sprintf("Bearer %s", secret)
	if len(secret) == 0 {
		panic("No Secret found")
	}
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if c.Request.URL.Path == "/v1/health" {
			return
		}
		if len(token) == 0 {
			c.AbortWithStatus(401)
			return
		}
		if token != formattedSecret {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
