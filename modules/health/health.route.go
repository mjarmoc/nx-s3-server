package health

import "github.com/gin-gonic/gin"

func NewHealth(c *HealthController, r *gin.RouterGroup) {

	r.GET("/health", c.Status)

}
