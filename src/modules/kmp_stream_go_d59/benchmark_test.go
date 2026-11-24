package mod_kmp_stream_go_d59

import "testing"

func Benchmark_KmpGo(b *testing.B) {
	txtB, patB := "abcabcdabxabcdabcdabde", "abcdabd"
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = KmpFindAll(txtB, patB)
	}
	_ = sink
}
