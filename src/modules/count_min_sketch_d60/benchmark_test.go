package mod_count_min_sketch_d60

import "testing"

func Benchmark_CMSketch(b *testing.B) {
	cm := NewCountMinSketch(4, 256)
	cm.Add(77, 1)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		cm.Add(uint64(i&63), 1)
	sink += cm.Estimate(uint64(i & 63))
	}
	_ = sink
}
