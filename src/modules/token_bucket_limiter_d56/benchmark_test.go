package mod_token_bucket_limiter_d56

import "testing"

func Benchmark_Limiter(b *testing.B) {
	tbB := NewTokenBucket(100.0, 50.0)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = tbB.Allow(float64(i)*0.01, 1.0)
	}
	_ = sink
}
