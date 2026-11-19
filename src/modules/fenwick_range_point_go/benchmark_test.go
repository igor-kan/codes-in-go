package mod_fenwick_range_point_go

import "testing"

func Benchmark_BitGo(b *testing.B) {
	f := NewFenwickTree(16); f.Add(3, 5); f.Add(9, 4)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(f.KthSmallest(6))
	}
	_ = sink
}
