package mod_radix_compact_trie_d56

import "testing"

func Benchmark_Trie(b *testing.B) {
	trB := NewPrefixTrie()
	trB.Insert("compute")
	trB.Insert("compile")
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = trB.CountPrefix("comp")
	}
	_ = sink
}
