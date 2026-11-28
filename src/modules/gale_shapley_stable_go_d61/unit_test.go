package mod_gale_shapley_stable_go_d61

import "testing"

func TestUnit_StableMarriage(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	pa := [][]int{{0, 1}, {1, 0}}
	pb := [][]int{{1, 0}, {0, 1}}
	got := StableMatch(pa, pb)
	if len(got) != 2 || got[0] != 0 || got[1] != 1 {
		t.Fatalf("unexpected matching: %v", got)
	}
	pa2 := [][]int{{1, 0}, {0, 1}}
	pb2 := [][]int{{1, 0}, {0, 1}}
	got2 := StableMatch(pa2, pb2)
	if got2[0] != 1 || got2[1] != 0 {
		t.Fatalf("unexpected cross matching: %v", got2)
	}
	pairs := MatchedPairs(got)
	if len(pairs) != 2 || pairs[0].A != 0 || pairs[0].B != 0 {
		t.Fatalf("bad pair projection: %v", pairs)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
