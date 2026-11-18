package mod_cuckoo_filter_fp8_go

import "testing"

func Benchmark_CuckooGo(b *testing.B) {
	cf := NewCuckooFilter(32); cf.Insert(42)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		if cf.Lookup(42) { sink++ }
	}
	_ = sink
}
