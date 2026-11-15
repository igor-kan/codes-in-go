package mod_lfu_cache_o1_buckets

import "testing"

func TestLFUMixedWorkloadInvariants(t *testing.T) {
	c, _ := NewLFUCache(8)
	for step := 0; step < 300; step++ {
		k := (step*13 + 7) % 19
		if step%3 == 0 {
			_, _ = c.Get(k)
		} else {
			c.Put(k, step)
		}
		if err := VerifyLFUInvariants(c); err != nil {
			t.Fatal(err)
		}
	}
}
