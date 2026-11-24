package mod_dijkstra_sp_go_d59

import "testing"

func TestUnit_Sssp(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	adj := [][]WeightedArc{{{1, 4}, {2, 1}}, {{3, 1}}, {{1, 2}, {3, 5}}, {}}
	d := DijkstraDist(4, 0, adj)
	if d[0] != 0 || d[1] != 3 || d[3] != 4 {
		t.Fatalf("bad dist: %v", d)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
