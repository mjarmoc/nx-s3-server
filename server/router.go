package server

import (
	"github.com/mjarmoc/nx-s3-server/modules/cache"
	"github.com/mjarmoc/nx-s3-server/modules/health"

	"github.com/gin-gonic/gin"
	"github.com/mjarmoc/nx-s3-server/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	health.NewHealth(new(health.HealthController), v1)
	cache.NewCache(new(cache.CacheController), v1)
	return router

}
