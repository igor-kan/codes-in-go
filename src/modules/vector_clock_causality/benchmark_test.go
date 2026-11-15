package mod_vector_clock_causality

import "testing"

func BenchmarkVectorClockCompare(b *testing.B) {
	x, _ := NewVectorClock(16)
	y, _ := NewVectorClock(16)
	_ = x.Tick(2)
	_ = y.Tick(7)
	for i := 0; i < b.N; i++ {
		_, _ = x.Compare(y)
	}
}
