package pack

import (
	"testing"
)

func BenchmarkPrintWeatherParallelism(b *testing.B) {
	b.SetParallelism(32)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			PrintWeather()
		}
	})
}
