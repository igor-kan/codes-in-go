package mod_trie_wildcard_d59

import "testing"

func Benchmark_Wtrie(b *testing.B) {
	wtB := NewWildTrie(); wtB.Add("search")
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = wtB.Search("s.a.ch")
	}
	_ = sink
}
