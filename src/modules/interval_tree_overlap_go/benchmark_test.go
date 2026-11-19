package mod_interval_tree_overlap_go

import "testing"

func Benchmark_IvTreeGo(b *testing.B) {
	tr := BuildIntervalTree([]ClosedIv{{1, 5, 1}, {3, 8, 2}, {10, 12, 3}})
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(len(tr.QueryOverlap(4, 9)))
	}
	_ = sink
}
