package mod_tarjan_scc_condense_go

import "testing"

func Benchmark_TarjanGo(b *testing.B) {
	es := [][2]int{{0, 1}, {1, 2}, {2, 0}}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(TarjanSCC(3, es).Count)
	}
	_ = sink
}
