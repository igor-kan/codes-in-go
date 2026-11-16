package mod_reservoir_quantiles_gk

import "testing"

func BenchmarkGKObserve(b *testing.B) {
	g, _ := NewGKSummary(0.01)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Observe(float64(i & 1023))
	}
}
