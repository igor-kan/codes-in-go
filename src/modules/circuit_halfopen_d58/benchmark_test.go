package mod_circuit_halfopen_d58

import "testing"

func Benchmark_Breaker(b *testing.B) {
	cbB := NewBreaker(5, 100)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = cbB.Allow(int64(i))
	}
	_ = sink
}
