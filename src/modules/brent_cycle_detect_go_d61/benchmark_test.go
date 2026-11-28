package mod_brent_cycle_detect_go_d61

import "testing"

func Benchmark_BrentCycle(b *testing.B) {
	f := func(x int) int { return (x*75 + 74) % 1021 }
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(BrentCycle(f, 17).Lambda)
	}
	_ = sink
}
