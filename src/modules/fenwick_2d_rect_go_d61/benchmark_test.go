package mod_fenwick_2d_rect_go_d61

import "testing"

func Benchmark_Fenwick2D(b *testing.B) {
	f := NewFenwick2D(32, 32)
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			f.Add(i, j, int64(i+j))
		}
	}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += f.RectSum(RectQuery{X1: 0, Y1: 0, X2: i & 31, Y2: 31})
	}
	_ = sink
}
