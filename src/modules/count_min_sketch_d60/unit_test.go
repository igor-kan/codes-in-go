package mod_count_min_sketch_d60

import "testing"

func TestUnit_CMSketch(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	cm := NewCountMinSketch(4, 128)
	cm.Add(101, 5)
	cm.Add(202, 2)
	cm.Add(101, 3)
	if cm.Estimate(101) < 8 || cm.Estimate(101) > 10 {
		t.Fatalf("bad est %d", cm.Estimate(101))
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
