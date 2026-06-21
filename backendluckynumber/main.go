package main

import (
	"backendluckynumber/function"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(function.CORSMiddleware())

	r.GET("/api/random-number", function.GetRandomNumberHandler)
	r.POST("/api/random-numbers", function.RandomNumbersHandler)

	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
