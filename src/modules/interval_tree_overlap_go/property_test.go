package mod_interval_tree_overlap_go

import "testing"

func TestProp_IvTreeGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	tr := BuildIntervalTree([]ClosedIv{{1, 3, 1}}); if len(tr.QueryOverlap(4, 5)) != 0 { t.Fatal("expected 0") }
}
