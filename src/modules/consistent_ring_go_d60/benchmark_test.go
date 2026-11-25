package mod_consistent_ring_go_d60

import "testing"

func Benchmark_HashRing(b *testing.B) {
	ring := NewConsistentRing(16)
	ring.AddNode("s1")
	ring.AddNode("s2")
	ring.AddNode("s3")
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		if nd, ok := ring.RouteKey("session-key"); ok { sink += int64(len(nd)) }
	}
	_ = sink
}
