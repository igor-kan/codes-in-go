package mod_crdt_lww_element_set

import "testing"

func BenchmarkLWWMerge(b *testing.B) {
	a := NewLWWSet(BiasRemoveWins)
	c := NewLWWSet(BiasRemoveWins)
	for i := 0; i < 64; i++ {
		k := string(rune('a' + (i % 26)))
		_ = a.Add(k, uint64(i))
		_ = c.Add(k, uint64(i+1))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cp := a.Clone()
		_ = cp.Merge(c)
	}
}
