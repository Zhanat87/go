package pack

import (
	"testing"
)

// все тестовые методы должны начинаться с 'Test'
func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Log("add numbers error")
		t.Fail()
	}

	result = Add(1, 2, 3, 4)
	if result != 10 {
		t.Error("add numbers error 2")
	}

	result = Add()
	if result != 0 {
		t.Error("add numbers error 3")
	}
}