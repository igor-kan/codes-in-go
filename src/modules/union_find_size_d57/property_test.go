package mod_union_find_size_d57

import "testing"

func TestProp_DsuSize(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	uf := NewUnionFindSize(4)
	_ = uf.Union(0, 1)
	if uf.Union(0, 1) {
		t.Fatalf("duplicate union returned true")
	}
}
