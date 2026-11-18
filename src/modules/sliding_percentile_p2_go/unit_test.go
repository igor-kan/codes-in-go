package mod_sliding_percentile_p2_go

import "testing"

func TestUnit_WindowQuant(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	w := NewSlidingQuantileWindow(5)
	for _, v := range []float64{100, 1, 5, 3, 9, 7} {
		w.Push(v)
	}
	s := w.Summary()
	if s.Min != 1 || s.Median != 5 || s.Max != 9 {
		t.Fatalf("unexpected summary: %+v", s)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
