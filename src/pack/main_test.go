package pack

import (
	"testing"
)

func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Log("add numbers error")
		t.Fail()
	}
}