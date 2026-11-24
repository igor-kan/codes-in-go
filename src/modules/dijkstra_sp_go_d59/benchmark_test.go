package mod_dijkstra_sp_go_d59

import "testing"

func Benchmark_Sssp(b *testing.B) {
	adjB := [][]WeightedArc{{{1, 2}, {2, 7}}, {{2, 3}}, {}}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = DijkstraDist(3, 0, adjB)
	}
	_ = sink
}
