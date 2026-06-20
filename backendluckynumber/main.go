package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RandomRequest struct {
	Count int `json:"count"`
}

// 1. ยกโค้ดดัก CORS เดิมของคุณมาทำเป็น Middleware ตรงนี้ (ไม่พึ่งแพ็กเกจอื่นเลย)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 2. แยก Handler ตัวที่มีปัญหาจุดแดงออกมา
func RandomNumbersHandler(c *gin.Context) {
	var req RandomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณากรอกข้อมูลให้ถูกต้อง"})
		return
	}

	// เงื่อนไข 1-10 ที่เราต้องการจะเทสให้จุดแดงหายไป
	if req.Count < 1 || req.Count > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "สามารถสุ่มได้ตั้งแต่ 1 ถึง 10 เบอร์เท่านั้น"})
		return
	}
	luckyList := []string{"0991234567", "0997654321"} // เปลี่ยนเป็นตัวแปรสุ่มจริงของคุณนะ

	// ต้องส่งกลับโดยมีคีย์ชื่อ "data" ล้อไปกับที่หน้าบ้านเขียน result.data ไว้ครับ!
	c.JSON(http.StatusOK, gin.H{
		"data": luckyList,
	})
}

// ⚠️ อย่าลืมก๊อปปี้ Logic การวนลูปสุ่มเบอร์แล้วส่งกลับ (บรรทัด 58 ของคุณ) มาแปะต่อตรงนี้ด้วยนะครับ

func GetRandomNumberHandler(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	prefix := "099"
	suffix := rand.Intn(9000000) + 1000000
	number := fmt.Sprintf("%s%d", prefix, suffix)
	c.JSON(http.StatusOK, gin.H{"phoneNumber": number})
}

// 3. ฟังก์ชันรวมกลุ่มเส้นทางเพื่อให้ไฟล์ Test เรียกใช้งานได้
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/api/random-number", GetRandomNumberHandler)
	r.POST("/api/random-numbers", RandomNumbersHandler)

	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
