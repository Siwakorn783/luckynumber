package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RandomRequest struct {
	Count int `json:"count" binding:"required"`
}

// ฟังก์ชันจำลองการสร้างเบอร์มงคล (ในระบบจริงอาจจะดึงมาจาก Database)
func generateLuckyNumbers(count int) []string {
	rand.Seed(time.Now().UnixNano())
	prefixes := []string{"088", "099", "061", "062", "095"}
	var results []string

	for i := 0; i < count; i++ {
		// สุ่มคำนำหน้า
		prefix := prefixes[rand.Intn(len(prefixes))]
		// สุ่มเลขที่เหลืออีก 7 หลัก
		suffix := fmt.Sprintf("%07d", rand.Intn(10000000))
		results = append(results, prefix+suffix)
	}
	return results
}

func main() {
	r := gin.Default()

	// ใช้ CORS เพื่อให้หน้าเว็บยิง API เข้ามาได้
	r.Use(cors.Default())

	// Endpoint นี้คือที่ที่หน้าเว็บจะกดปุ่มมาขอเลข
	r.GET("/api/random-number", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		prefix := "099"
		suffix := rand.Intn(9000000) + 1000000 // สุ่มเลข 7 หลัก
		number := fmt.Sprintf("%s%d", prefix, suffix)

		c.JSON(200, gin.H{"phoneNumber": number})
	})

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/api/random-numbers", func(c *gin.Context) {
		var req RandomRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณากรอกข้อมูลให้ถูกต้อง"})
			return
		}

		// ตรวจสอบเงื่อนไข 1-10 เท่านั้น
		if req.Count < 1 || req.Count > 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "สามารถสุ่มได้ตั้งแต่ 1 ถึง 10 เบอร์เท่านั้น"})
			return
		}

		// สุ่มเบอร์และส่งกลับ
		luckyNumbers := generateLuckyNumbers(req.Count)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    luckyNumbers,
		})
	})
	r.Run(":8080")

}
