package mod_count_min_sketch_d57

import "testing"

func TestUnit_Cms(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	cm := NewCountMin(64)
	cm.Add(42, 7)
	cm.Add(99, 3)
	if cm.Query(42) < 7 || cm.Query(99) < 3 {
		t.Fatalf("count-min underestimated")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
