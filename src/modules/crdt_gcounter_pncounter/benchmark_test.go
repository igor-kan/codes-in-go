package mod_crdt_gcounter_pncounter

import "testing"

func BenchmarkCRDTMerge(b *testing.B) {
	x, _ := NewPNCounter(8)
	y, _ := NewPNCounter(8)
	_ = y.Add(3, 42)
	for i := 0; i < b.N; i++ {
		_ = x.Merge(y)
	}
}
