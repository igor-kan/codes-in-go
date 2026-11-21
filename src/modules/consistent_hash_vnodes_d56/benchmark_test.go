package mod_consistent_hash_vnodes_d56

import "testing"

func Benchmark_Ring(b *testing.B) {
	hrB := NewHashRing([]string{"n1", "n2", "n3"}, 12)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = hrB.Lookup("hot-key")
	}
	_ = sink
}
