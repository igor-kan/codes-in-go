package mod_kmp_prefix_scan_d60

import "testing"

func Benchmark_KmpMatcher(b *testing.B) {
	txt := "ababababcababababc"
	pat := "ababc"
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(len(KmpFindAll(txt, pat)))
	}
	_ = sink
}
