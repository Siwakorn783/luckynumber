package lucky

import (
	"testing"
)

func TestGenerateLuckyNumbers(t *testing.T) {

	count := 3

	results := GenerateLuckyNumbers(count)

	if len(results) != count {
		t.Errorf("ต้องได้จำนวนผลลัพท์%dแต่ได้%d", count, len(results))

	}
	for _, num := range results {
		if len(num) != 10 {
			t.Errorf("เบอร์ต้องมี%dหลักแต่ได้%d", 10, len(num))
		}
	}
}
