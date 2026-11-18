package mod_radix_patricia_router_go

import "testing"

func Benchmark_RadixRouter(b *testing.B) {
	rt := NewBytePrefixRouter(); rt.Insert("/api/v1", 7)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(rt.MatchLongest("/api/v1/items").Handler)
	}
	_ = sink
}
