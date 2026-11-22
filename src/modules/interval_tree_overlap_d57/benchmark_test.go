package mod_interval_tree_overlap_d57

import "testing"

func Benchmark_Sweep(b *testing.B) {
	spB := []TimeSpan{{0, 10}, {3, 7}, {5, 12}, {8, 15}}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = MaxIntervalOverlap(spB)
	}
	_ = sink
}
