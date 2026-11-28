package mod_boruvka_mst_go_d61

import "testing"

func Benchmark_BoruvkaMst(b *testing.B) {
	n := 32
	edges := make([]WeightedEdge, 0, 64)
	for i := 0; i < n; i++ {
		edges = append(edges, WeightedEdge{U: i, V: (i*7 + 3) % n, W: int64((i*13)%97 + 1)})
		edges = append(edges, WeightedEdge{U: i, V: (i + 1) % n, W: int64((i*5)%41 + 2)})
	}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += BoruvkaMST(n, edges).Weight
	}
	_ = sink
}
