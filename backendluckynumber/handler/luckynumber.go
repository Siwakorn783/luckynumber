package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// --- Helper Function (ใช้ร่วมกัน) ---

// generateLuckyNumber เป็นฟังก์ชันกลางสำหรับการสุ่มเลข
func generateLuckyNumber(prefix string) string {
	// สุ่มเลข 7 หลัก (เพื่อให้รวมกับ prefix 3 หลักเป็น 10 หลัก)
	suffix := rand.Intn(10000000)
	return fmt.Sprintf("%s%07d", prefix, suffix)
}

// --- Handlers ---

func RandomNumbersHandler(c *gin.Context) {
	var req struct {
		Count int `json:"count"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลที่ส่งมาไม่ถูกต้อง"})
		return
	}

	if req.Count < 1 || req.Count > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "สามารถสุ่มได้ตั้งแต่ 1 ถึง 10 เบอร์เท่านั้น"})
		return
	}

	// รายการ prefix ที่ต้องการให้สุ่ม
	prefixes := []string{"081", "089", "099", "062", "063", "091"}

	var luckyList []string
	for i := 0; i < req.Count; i++ {
		// สุ่มเลือก index จาก slice ของ prefixes
		randomPrefix := prefixes[rand.Intn(len(prefixes))]

		// เรียกใช้ฟังก์ชันเดิมโดยส่ง prefix ที่สุ่มได้เข้าไป
		luckyList = append(luckyList, generateLuckyNumber(randomPrefix))
	}

	c.JSON(http.StatusOK, gin.H{"data": luckyList})
}

func GetRandomoneNumberHandler(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	prefixes := []string{"088", "099", "061", "062", "095"}

	// สุ่ม prefix จากรายการ
	prefix := prefixes[rand.Intn(len(prefixes))]

	// ใช้ฟังก์ชันกลางที่สร้างขึ้น
	number := generateLuckyNumber(prefix)

	c.JSON(http.StatusOK, gin.H{"phoneNumber": number})
}
