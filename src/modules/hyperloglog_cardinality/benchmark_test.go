package mod_hyperloglog_cardinality

import "testing"

func BenchmarkHLLAdd(b *testing.B) {
	h, _ := NewHLL(12, 1)
	for i := 0; i < b.N; i++ {
		h.Add(uint64(i))
	}
}
