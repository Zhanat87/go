package poms

import (
	"net/http"
	"os"
	"poms/ctrl"
	"testing"
	"fmt"
)

// запускается вначале перед всеми тестами
func TestMain(m *testing.M) {
	ctrl.Setup()

	go http.ListenAndServe(":3000", new(GZipServer))

	// запускает все тесты
	fmt.Println("before run all test")
	m.Run()
	fmt.Println("after run all test")

	os.Exit(0)
}

func BenchmarkGetPurchaseOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:3000/api/purchaseOrders/1")
	}
}

func BenchmarkGetCurrencies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:3000/api/currencies")
	}
}
