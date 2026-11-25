package mod_hyperloglog_card_d60

import "testing"

func Benchmark_HllEstimator(b *testing.B) {
	hll := NewHyperLogLog(6)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		hll.Observe(uint64(i))
	sink += int64(hll.regs[i&63])
	}
	_ = sink
}
