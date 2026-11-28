package mod_douglas_peucker_go_d61

import "testing"

func Benchmark_PolylineRdp(b *testing.B) {
	pts := make([]Point, 0, 256)
	for i := 0; i < 256; i++ {
		pts = append(pts, Point{X: float64(i), Y: float64((i*i*7)%23)})
	}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(len(SimplifyPolyline(pts, 0.5)))
	}
	_ = sink
}
