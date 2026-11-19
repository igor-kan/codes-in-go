package mod_disjoint_set_union_go

import "testing"

func Benchmark_DsuGo(b *testing.B) {
	d := NewDisjointSet(8)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		d.Union(i&7, (i+1)&7); sink += int64(d.Components())
	}
	_ = sink
}
