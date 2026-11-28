package mod_gale_shapley_stable_go_d61

import "testing"

func Benchmark_StableMarriage(b *testing.B) {
	pa := [][]int{{0, 1, 2, 3}, {1, 0, 2, 3}, {2, 1, 0, 3}, {0, 2, 1, 3}}
	pb := [][]int{{3, 1, 2, 0}, {1, 3, 0, 2}, {2, 0, 3, 1}, {0, 1, 2, 3}}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(StableMatch(pa, pb)[0])
	}
	_ = sink
}
