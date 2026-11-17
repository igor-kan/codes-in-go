package mod_consistent_bounded_load

import "testing"

func BenchmarkBoundedAcquireRelease(b *testing.B) {
	ring, _ := NewBoundedHashRing([]string{"a","b","c"}, 8, 1.25)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		s, _ := ring.Acquire("k"); _ = ring.Release(s); sink++
	}
	_ = sink
}
