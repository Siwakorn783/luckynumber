package function

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
	// 1. สร้าง struct เพื่อรับค่าจาก JSON
	var req struct {
		Count int `json:"count"`
	}

	// 2. ใช้ ShouldBindJSON แทน c.Query
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลที่ส่งมาไม่ถูกต้อง"})
		return
	}

	// 3. ใช้ req.Count แทน
	if req.Count < 1 || req.Count > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "สามารถสุ่มได้ตั้งแต่ 1 ถึง 10 เบอร์เท่านั้น"})
		return
	}

	// ... (ส่วนที่เหลือของโค้ดคุณ เหมือนเดิม)
	var luckyList []string
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < req.Count; i++ {
		prefix := "099"
		suffix := rand.Intn(9000000) + 1000000
		number := fmt.Sprintf("%s%d", prefix, suffix)
		luckyList = append(luckyList, number)
	}

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
