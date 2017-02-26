package pack

import (
	"testing"
	"time"
)

// все тестовые методы должны начинаться с 'Test'
func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Log("add numbers error")
		t.Fail()
	}
	// go test pack -timeout 2s
	time.Sleep(3 * time.Second)

	result = Add(1, 2, 3, 4)
	if result != 10 {
		t.Error("add numbers error 2")
	}

	result = Add()
	if result != 0 {
		t.Error("add numbers error 3")
	}
}

func TestCanSubtractNumbers(t *testing.T) {
	result := Subtract(81, 1, 2, 82)
	if result != -4 {
		t.Log("subtract numbers error")
		t.Fail()
	}

	result = Subtract(2, 1)
	if result != 1 {
		t.Error("subtract numbers error 2")
	}

	result = Subtract(2)
	if result != 2 {
		t.Error("subtract numbers error 3")
	}
}
