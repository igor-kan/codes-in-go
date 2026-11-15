package mod_timing_wheel_scheduler

import "testing"

func BenchmarkTimingWheelCycle(b *testing.B) {
	tw, _ := NewTimingWheel(64)
	for i := 0; i < b.N; i++ {
		_, _ = tw.Schedule(uint64(i+1), uint64((i%128)+1), "x")
		_ = tw.Advance(1)
	}
}
