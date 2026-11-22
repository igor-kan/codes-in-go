package mod_union_find_size_d57

import "testing"

func Benchmark_DsuSize(b *testing.B) {
	ufB := NewUnionFindSize(16)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		ufB.Union(i&15, (i+1)&15)
	}
	_ = sink
}
