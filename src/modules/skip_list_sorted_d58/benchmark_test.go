package mod_skip_list_sorted_d58

import "testing"

func Benchmark_SkipChunk(b *testing.B) {
	idxB := NewSortedIndex([]int64{2, 5, 8, 12, 16, 20, 25})
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = idxB.CountRange(5, 20)
	}
	_ = sink
}
