package mod_interval_tree_overlap_d57

import "testing"

func TestProp_Sweep(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	disjoint := []TimeSpan{{0, 2}, {2, 4}, {4, 6}}
	if g := MaxIntervalOverlap(disjoint); g != 1 {
		t.Fatalf("half-open adjacency overlapped: %d", g)
	}
}
