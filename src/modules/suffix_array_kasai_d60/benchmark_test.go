package mod_suffix_array_kasai_d60

import "testing"

func Benchmark_SuffixIndex(b *testing.B) {
	sample := "banana$bandana"
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sa, _ := BuildSuffixAndLCP(sample)
	sink += int64(sa[0])
	}
	_ = sink
}
