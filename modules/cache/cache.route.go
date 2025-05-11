package cache

import "github.com/gin-gonic/gin"

func NewCache(c *CacheController, r *gin.RouterGroup) {
	cacheGroup := r.Group("cache")
	cacheGroup.GET("/", c.List)
	cacheGroup.GET("/:hash", c.Retrieve)
	cacheGroup.PUT("/:hash", c.Save)
}
