package mod_fenwick_2d_rect_go_d61

import "testing"

func TestUnit_Fenwick2D(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	f := NewFenwick2D(4, 4)
	f.Add(1, 1, 5)
	f.Add(3, 3, 7)
	f.Add(1, 3, -2)
	if got := f.RectSum(RectQuery{X1: 0, Y1: 0, X2: 3, Y2: 3}); got != 10 {
		t.Fatalf("full sum = %d, want 10", got)
	}
	if got := f.RectSum(RectQuery{X1: 1, Y1: 1, X2: 1, Y2: 1}); got != 5 {
		t.Fatalf("single cell = %d, want 5", got)
	}
	if got := f.RectSum(RectQuery{X1: 1, Y1: 2, X2: 3, Y2: 3}); got != 5 {
		t.Fatalf("corner sum = %d, want 5", got)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
