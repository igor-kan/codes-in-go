package mod_bipartite_kuhn_match_d56

import "testing"

func TestUnit_Kuhn(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	adj := [][]int{{0, 1}, {0}, {1, 2}, {2}}
	if m := MaxBipartiteMatch(4, 3, adj); m != 3 {
		t.Fatalf("expected 3, got %d", m)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
