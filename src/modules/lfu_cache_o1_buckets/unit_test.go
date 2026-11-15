package mod_lfu_cache_o1_buckets

import "testing"

func TestLFUEvictionAndTieBreak(t *testing.T) {
	c, _ := NewLFUCache(2)
	c.Put(1, 10)
	c.Put(2, 20)
	if v, ok := c.Get(1); !ok || v != 10 {
		t.Fatal("expected key 1 hit")
	}
	c.Put(3, 30) // evicts key 2 (freq=1), keeps key 1 (freq=2)
	if _, ok := c.Get(2); ok {
		t.Fatal("key 2 should have been evicted")
	}
	if _, ok := c.Get(1); !ok {
		t.Fatal("key 1 should still be cached")
	}
	if err := VerifyLFUInvariants(c); err != nil {
		t.Fatal(err)
	}
}
