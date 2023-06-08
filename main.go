package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samarqandi/blog-gin/infrastructure"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		infrastructure.LoadEnv()
		infrastructure.NewDatabase()
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello, Gin!"})
	})
	router.Run(":8000")
}
