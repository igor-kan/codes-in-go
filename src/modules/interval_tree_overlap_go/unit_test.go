package mod_interval_tree_overlap_go

import "testing"

func TestUnit_IvTreeGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	tr := BuildIntervalTree([]ClosedIv{{1, 5, 1}, {10, 15, 2}, {4, 12, 3}, {20, 25, 4}})
	hits := tr.QueryOverlap(6, 11)
	if len(hits) != 2 {
		t.Fatalf("expected 2 overlapping intervals, got %v", hits)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
