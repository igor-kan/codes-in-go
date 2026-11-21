package mod_bipartite_kuhn_match_d56

import "testing"

func TestProp_Kuhn(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	adj := [][]int{{0}, {0}, {0}}
	if m := MaxBipartiteMatch(3, 1, adj); m != 1 {
		t.Fatalf("bottleneck violated: %d", m)
	}
}
