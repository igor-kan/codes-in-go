package mod_count_min_sketch_d57

import "testing"

func Benchmark_Cms(b *testing.B) {
	cmB := NewCountMin(64); cmB.Add(1, 5)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		cmB.Add(uint64(i&15), 1); _ = cmB.Query(1)
	}
	_ = sink
}
