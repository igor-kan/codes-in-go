package mod_interval_tree_overlap_d57

import "testing"

func TestUnit_Sweep(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	sp := []TimeSpan{{1, 5}, {2, 6}, {4, 8}, {5, 9}}
	if g := MaxIntervalOverlap(sp); g != 3 {
		t.Fatalf("expected 3, got %d", g)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
