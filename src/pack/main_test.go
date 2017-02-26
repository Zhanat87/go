package pack

import (
	"testing"
	//"time"
)

// все тестовые методы должны начинаться с 'Test'
func TestCanAddNumbers(t *testing.T) {
	// go test pack -v -short
	if testing.Short() {
		t.Skip("Skipping long tests")
	}
	result := Add(1, 2)
	if result != 3 {
		t.Log("add numbers error")
		t.FailNow()
	}
	// go test pack -timeout 2s
	//time.Sleep(3 * time.Second)

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
	result := Subtract(2, 1)
	if result != 1 {
		t.Fatal("subtract numbers error 2")
	}

	result = Subtract(81, 1, 2, 82)
	if result != -4 {
		t.Log("subtract numbers error")
		t.Fail()
	}

	result = Subtract(2)
	if result != 2 {
		t.Error("subtract numbers error 3")
	}
}

func TestCanMultiplyNumbers(t *testing.T) {
	result := Multiply()
	if result != 0 {
		t.Fatal("multiply numbers error")
	}

	result = Multiply(1, 2, 3, 4, 5, 6, 7)
	if result != 5040 {
		t.Fatal("multiply numbers error 2")
	}
}

func TestCanDivideNumbers(t *testing.T) {
	// go test pack -v
	if testing.Verbose() {
		t.Skip("Not implemented yet")
	}
	//result := Divide()
	//if result != 0 {
	//	t.Fatal("divide numbers error")
	//}
	//
	//result = Multiply(1, 2, 3, 4, 5, 6, 7)
	//if result != 5040 {
	//	t.Fatal("divide numbers error 2")
	//}
}
