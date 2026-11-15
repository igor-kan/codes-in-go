package mod_union_find_weighted

import "testing"

func BenchmarkWeightedDSURelate(b *testing.B) {
	d, _ := NewWeightedDSU(1024)
	for i := 0; i < b.N; i++ {
		_, _ = d.Relate(i&1023, (i+1)&1023, 1.0)
	}
}
