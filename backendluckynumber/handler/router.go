package handler

import "github.com/gin-gonic/gin"

// SetupRouter ทำหน้าที่รวม Routes ทั้งหมด
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/random-number", GetRandomoneNumberHandler)
	r.POST("/api/random-numbers", RandomNumbersHandler)
	return r
}
