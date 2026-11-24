package mod_aho_multi_scan_d59

import "testing"

func Benchmark_AcGo(b *testing.B) {
	acB := BuildAc([]string{"go", "fast", "alg"})
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = acB.CountHits("gofastalgorithmgo")
	}
	_ = sink
}
