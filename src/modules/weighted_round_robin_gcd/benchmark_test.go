package mod_weighted_round_robin_gcd

import "testing"

func Benchmark_SmoothWrr(b *testing.B) {
	wrr := NewSmoothWRR([]int{3, 2, 1})
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(wrr.NextPeer())
	}
	_ = sink
}
