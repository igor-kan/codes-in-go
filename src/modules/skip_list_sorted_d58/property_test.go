package mod_skip_list_sorted_d58

import "testing"

func TestProp_SkipChunk(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	idx := NewSortedIndex([]int64{1, 2, 3, 4})
	if idx.CountRange(10, 20) != 0 || idx.CountRange(0, 5) != 4 {
		t.Fatalf("boundary error")
	}
}
