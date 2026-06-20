package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
func TestRandomNumbersHandler_TableDriven(t *testing.T) {
	router := SetupRouter()

	// 1. สร้างตารางโครงสร้างข้อมูล (Define the Test Case Struct)
	// เพื่อกำหนดว่าใน 1 แถวของตาราง เราจะบันทึกข้อมูลอะไรสำหรับใช้เทสบ้าง
	tests := []struct {
		name           string // ชื่อสั้นๆ เพื่อบอกว่าเคสนี้กำลังเทสอะไร
		inputBody      string // ข้อมูล JSON ปลอมที่จะส่งเข้า API
		expectedStatus int    // รหัส Status Code ที่เราคาดหวังว่าจะได้กลับมา
		expectedBody   string // ข้อความสำคัญที่เราคาดหวังว่าจะเจอในผลลัพธ์
	}{
		// 2. เติมข้อมูลลงตาราง (The actual rows of data)
		{
			name:           "ส่งค่าเกินขอบเขต (เลข 11 ต้องโดนบล็อก)",
			inputBody:      `{"count": 11}`,
			expectedStatus: http.StatusBadRequest, // 400
			expectedBody:   "สามารถสุ่มได้ตั้งแต่ 1 ถึง 10 เบอร์เท่านั้น",
		},
		{
			name:           "ส่งค่าน้อยกว่าขอบเขต (เลข 0 ต้องโดนบล็อก)",
			inputBody:      `{"count": 0}`,
			expectedStatus: http.StatusBadRequest, // 400
			expectedBody:   "สามารถสุ่มได้ตั้งแต่ 1 ถึง 10 เบอร์เท่านั้น",
		},
		{
			name:           "ส่งค่าถูกต้อง (เลข 5 ต้องผ่านฉลุย)",
			inputBody:      `{"count": 5}`,
			expectedStatus: http.StatusOK, // 200
			expectedBody:   `"data"`,      // ต้องมีคีย์ data โผล่มาให้หน้าบ้าน
		},
		{
			name:           "ส่งข้อมูลพังๆ ไม่ใช่ JSON (ต้องโดนดักตั้งแต่ประตูแรก)",
			inputBody:      `{"count": "ไม่ใช่ตัวเลข"}`,
			expectedStatus: http.StatusBadRequest, // 400
			expectedBody:   "กรุณากรอกข้อมูลให้ถูกต้อง",
		},
	}

	// 3. วนลูปเพื่อรันเทสทีละแถวในตาราง (Iterate over the table)
	for _, tc := range tests {
		// t.Run จะแยกข้อย่อยให้ตอนรัน ทำให้เรารู้ว่าแถวไหนผ่าน แถวไหนพัง
		t.Run(tc.name, func(t *testing.T) {
			// Arrange: ดึงข้อมูลจากแถวปัจจุบัน (tc) มาตั้งค่า Request ปลอม
			req, _ := http.NewRequest("POST", "/api/random-numbers", bytes.NewBufferString(tc.inputBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Act: ยิงข้อมูลเข้า Handler
			router.ServeHTTP(w, req)

			// Assert: ตรวจสอบ Status Code ประจำแถว
			if w.Code != tc.expectedStatus {
				t.Errorf("เคส [%s] พัง: คาดหวัง Status %d แต่ระบบดันได้ %d", tc.name, tc.expectedStatus, w.Code)
			}

			// Assert: ตรวจสอบข้อความแจ้งเตือนประจำแถว
			if !strings.Contains(w.Body.String(), tc.expectedBody) {
				t.Errorf("เคส [%s] พัง: คาดหวังว่าจะเจอคำว่า '%s' แต่ในเนื้อหาจริงกลับไม่มี", tc.name, tc.expectedBody)
			}
		})
	}
}
