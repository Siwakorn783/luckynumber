package main

import (
	"backendluckynumber/handler"
	"backendluckynumber/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/api/random-number", handler.GetRandomoneNumberHandler)
	r.POST("/api/random-numbers", handler.RandomNumbersHandler)

	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
