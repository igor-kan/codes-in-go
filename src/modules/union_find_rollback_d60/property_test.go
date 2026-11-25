package mod_union_find_rollback_d60

import "testing"

func TestProp_RollbackDsu(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	d := NewDsuWithRollback(6)
	s0 := d.Snapshot()
	d.Union(0, 1); d.Union(2, 3); d.Union(0, 3)
	d.Rollback(s0)
	if d.Components() != 6 {
		t.Fatal("full rollback did not restore 6 components")
	}
}
