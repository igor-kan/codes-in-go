package mod_bloom_double_hash_d56

import "testing"

func Benchmark_Bloom(b *testing.B) {
	bfB := NewBloomFilter(512, 4); bfB.Add(99)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = bfB.MightContain(uint64(i))
	}
	_ = sink
}
