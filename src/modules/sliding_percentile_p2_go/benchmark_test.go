package mod_sliding_percentile_p2_go

import "testing"

func Benchmark_WindowQuant(b *testing.B) {
	w := NewSlidingQuantileWindow(5); w.Push(1); w.Push(2); w.Push(3)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(w.Summary().Median)
	}
	_ = sink
}
