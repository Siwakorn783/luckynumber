package lucky

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateLuckyNumbers(count int) []string {
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
