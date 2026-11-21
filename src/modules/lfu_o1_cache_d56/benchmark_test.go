package mod_lfu_o1_cache_d56

import "testing"

func Benchmark_Lfu(b *testing.B) {
	cacheB := NewLfuCache(8)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		cacheB.Put(i&15, i); _, _ = cacheB.Get(i & 7)
	}
	_ = sink
}
