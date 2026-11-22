package mod_hyperloglog_sketch_d57

import "testing"

func Benchmark_Hll(b *testing.B) {
	skB := NewHll64()
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		skB.Add(uint64(i))
	}
	_ = sink
}
