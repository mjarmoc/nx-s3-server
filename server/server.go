package server

import (
	"fmt"
	"net/http"

	"github.com/mjarmoc/nx-s3-server/config"

	"github.com/gin-gonic/gin"
)

func Init() {
	config := config.GetConfig()
	address := config.GetString("server.address")
	r := NewRouter()
	r.Use(ErrorHandler())
	r.Run(address)
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				fmt.Println("Recovered from panic:", err)

				// Return an error response to the client
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Recovered Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}
