package mod_aho_corasick_go

import "testing"

func BenchmarkAcScan(b *testing.B) {
	ac, _ := CompilePatterns([][]byte{[]byte("alpha"), []byte("beta")}); doc := []byte("alpha_beta_alpha_gamma")
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(ac.CountMatches(doc))
	}
	_ = sink
}
