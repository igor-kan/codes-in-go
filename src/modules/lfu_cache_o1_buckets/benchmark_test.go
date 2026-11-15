package mod_lfu_cache_o1_buckets

import "testing"

func BenchmarkLFUPutGet(b *testing.B) {
	c, _ := NewLFUCache(512)
	for i := 0; i < b.N; i++ {
		c.Put(i&1023, i)
		_, _ = c.Get(i & 255)
	}
}
