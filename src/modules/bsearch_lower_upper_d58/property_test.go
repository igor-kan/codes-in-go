package mod_bsearch_lower_upper_d58

import "testing"

func TestProp_BSearch(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	arr := []int64{2, 4, 6}
	for _, v := range arr {
		if CountSortedOccurrences(arr, v) != 1 {
			t.Fatalf("expected unique 1")
		}
	}
}
