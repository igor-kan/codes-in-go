package mod_bplus_leaf_cursor_go

import "testing"

func TestProp_BPlusCursor(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	idx := NewBPlusLeafIndex()
	idx.Put(5, 50); idx.Put(5, 99)
	res := idx.Scan(0, 10)
	if len(res) != 1 || res[0].Val != 99 {
		t.Fatal("upsert failed")
	}
}
