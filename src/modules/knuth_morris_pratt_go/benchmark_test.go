package mod_knuth_morris_pratt_go

import "testing"

func Benchmark_KmpGo(b *testing.B) {
	
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(len(KmpSearchAll("ababaababa", "aba")))
	}
	_ = sink
}
