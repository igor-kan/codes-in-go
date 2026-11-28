package mod_douglas_peucker_go_d61

import "testing"

func TestProp_PolylineRdp(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	base := make([]Point, 0, 50)
	for i := 0; i < 50; i++ {
		base = append(base, Point{X: float64(i), Y: 2 * float64(i)})
	}
	collinear := SimplifyPolyline(base, 0.001)
	if len(collinear) != 2 {
		t.Fatalf("collinear polyline should collapse to endpoints, got %d", len(collinear))
	}
	curve := make([]Point, 0, 64)
	for i := 0; i < 64; i++ {
		curve = append(curve, Point{X: float64(i), Y: float64((i*i)%17) / 4.0})
	}
	out := SimplifyPolyline(curve, 0.25)
	if len(out) < 2 || out[0] != curve[0] || out[len(out)-1] != curve[len(curve)-1] {
		t.Fatal("endpoints must be preserved")
	}
	if len(out) > len(curve) {
		t.Fatal("simplification must not add points")
	}
	if again := SimplifyPolyline(out, 0.25); len(again) > len(out) {
		t.Fatalf("simplification should not grow: %d vs %d", len(again), len(out))
	}
}
