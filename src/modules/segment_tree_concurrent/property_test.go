package mod_segment_tree_concurrent

import "testing"

func TestSegTreeParallelQueries(t *testing.T) {
	raw := []int64{10, -5, 7, 12, 0, 8, -3, 15}
	st, _ := NewConcurrentSegTree(raw)
	outs := parallelRangeQueries(st, [][2]int{{0, 4}, {2, 7}, {0, 8}})
	if outs[0] != fastSliceStats(raw, 0, 4) || outs[2] != fastSliceStats(raw, 0, 8) {
		t.Fatal("parallel segtree query mismatch")
	}
}
