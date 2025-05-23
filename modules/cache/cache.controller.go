package cache

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	s3 "github.com/mjarmoc/nx-s3-server/modules/aws"

	"github.com/gin-gonic/gin"
)

type CacheController struct{}

func (c CacheController) List(ctx *gin.Context) {
	svc := s3.GetService()
	files, err := svc.List(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, files)
}

func (c CacheController) Retrieve(ctx *gin.Context) {
	hash := ctx.Param("hash")
	svc := s3.GetService()
	file, err := svc.Get(ctx.Request.Context(), hash)
	if err != nil {
		switch err.(type) {
		case *s3.CacheNotFoundError:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.Data(http.StatusOK, "application/octet-stream", *file)
}

func (c CacheController) Save(ctx *gin.Context) {
	hash := ctx.Param("hash")
	svc := s3.GetService()
	fmt.Println("Start uploading")
	b, err := io.ReadAll(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	reader := bytes.NewReader(b)
	output, err := svc.Upload(ctx.Request.Context(), hash, reader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, output)
}
