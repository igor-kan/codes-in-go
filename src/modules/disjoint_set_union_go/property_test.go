package mod_disjoint_set_union_go

import "testing"

func TestProp_DsuGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	d := NewDisjointSet(4); if d.Components() != 4 { t.Fatal("expected 4 components") }
}
