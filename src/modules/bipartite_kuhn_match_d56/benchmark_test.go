package mod_bipartite_kuhn_match_d56

import "testing"

func Benchmark_Kuhn(b *testing.B) {
	adjB := [][]int{{0, 1}, {1, 2}, {0, 2}}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = MaxBipartiteMatch(3, 3, adjB)
	}
	_ = sink
}
