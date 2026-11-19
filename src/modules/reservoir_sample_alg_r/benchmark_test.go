package mod_reservoir_sample_alg_r

import "testing"

func Benchmark_ReservoirGo(b *testing.B) {
	rs := NewReservoirSampler(8, 7)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		rs.Offer(int64(i)); sink += rs.Items()[0]
	}
	_ = sink
}
