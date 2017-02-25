package pack

import (
	"bytes"
	"testing"
	"text/template"
)

// все тестовые методы для benchmark должны начинаться с 'Benchmark'
func BenchmarkExample(b *testing.B) {
	temp, _ := template.New("").Parse("Hello, Go")
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		temp.Execute(&buf, nil)
		buf.Reset()
	}
}
