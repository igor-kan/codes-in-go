package mod_bsearch_lower_upper_d58

import "testing"

func Benchmark_BSearch(b *testing.B) {
	arrB := []int64{1, 2, 2, 2, 4, 7, 9}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = CountSortedOccurrences(arrB, 2)
	}
	_ = sink
}
