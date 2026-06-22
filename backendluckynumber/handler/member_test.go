package handler

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
			expectedBody:   `"data"`,
		},
		{
			name:           "ส่งข้อมูลพังๆ ไม่ใช่ JSON (ต้องโดนดักตั้งแต่ประตูแรก)",
			inputBody:      `{"count": "ไม่ใช่ตัวเลข"}`,
			expectedStatus: http.StatusBadRequest, // 400
			expectedBody:   "ข้อมูลที่ส่งมาไม่ถูกต้อง",
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req, _ := http.NewRequest("POST", "/api/random-numbers", bytes.NewBufferString(tc.inputBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("เคส [%s] พัง: คาดหวัง Status %d แต่ระบบดันได้ %d", tc.name, tc.expectedStatus, w.Code)
			}

			if !strings.Contains(w.Body.String(), tc.expectedBody) {
				t.Errorf("เคส [%s] พัง: คาดหวังว่าจะเจอคำว่า '%s' แต่ในเนื้อหาจริงกลับไม่มี", tc.name, tc.expectedBody)
			}
		})
	}
}

func TestGenerateLuckyNumber(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		expected int
	}{
		{
			name:     "prefix 3 digits",
			prefix:   "123",
			expected: 10,
		},
		{
			name:     "prefix 1 digit",
			prefix:   "9",
			expected: 10,
		},
		{
			name:     "empty prefix",
			prefix:   "",
			expected: 7, // ถ้าไม่มี prefix ควรเหลือแค่ 7 หลักตาม rand.Intn
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateLuckyNumber(tt.prefix)
			if len(got) != tt.expected {
				t.Errorf("generateLuckyNumber() = %v (len: %d), want length %d", got, len(got), tt.expected)
			}
		})
	}
}
