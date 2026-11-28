package mod_douglas_peucker_go_d61

import "testing"

func TestUnit_PolylineRdp(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	line := []Point{{X: 0, Y: 0}, {X: 1, Y: 0.1}, {X: 2, Y: 0}, {X: 3, Y: 1}, {X: 4, Y: 0}}
	simp := SimplifyPolyline(line, 0.7)
	if len(simp) != 3 || simp[0].X != 0 || simp[1].X != 3 || simp[2].X != 4 {
		t.Fatalf("unexpected simplification: %v", simp)
	}
	if got := PerpendicularDistance(Point{X: 0, Y: 2}, Point{X: 0, Y: 0}, Point{X: 0, Y: 0}); got != 2 {
		t.Fatalf("degenerate distance = %v, want 2", got)
	}
	dup := []Point{{X: 1, Y: 1}, {X: 1, Y: 1}}
	if out := SimplifyPolyline(dup, 0.1); len(out) != 2 {
		t.Fatal("short inputs must be returned unchanged")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
