package mod_dijkstra_sp_go_d59

import "testing"

func TestProp_Sssp(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	adj := [][]WeightedArc{{{1, 5}}, {}}
	d := DijkstraDist(2, 1, adj)
	if d[0] < (1 << 59) {
		t.Fatalf("unreachable reached")
	}
}
