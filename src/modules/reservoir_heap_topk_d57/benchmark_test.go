package mod_reservoir_heap_topk_d57

import "testing"

func Benchmark_TopK(b *testing.B) {
	tkB := NewTopK(4)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		tkB.Push(int64(i * 7 % 97))
	}
	_ = sink
}
