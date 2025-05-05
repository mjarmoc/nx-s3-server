package cache

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type CacheController struct{}

func (c CacheController) Retrieve(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GET CACHE!")
	return
}

func (c CacheController) Save(ctx *gin.Context) {
	ctx.String(http.StatusOK, "SAVE CACHE!")
}