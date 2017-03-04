package pack

import (
	"testing"
)

func BenchmarkTemplateReportallocs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := NewUserReportallocs("test2")
		SayHelloReportallocs(user)
	}
}

func BenchmarkTemplateCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := NewUserReportallocs("test2")
		SayHelloReportallocs(user)
	}
}
