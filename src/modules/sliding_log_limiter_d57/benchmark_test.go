package mod_sliding_log_limiter_d57

import "testing"

func Benchmark_SLog(b *testing.B) {
	slB := NewSlidingLog(16, 100)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = slB.Allow(int64(i * 10))
	}
	_ = sink
}
