package mod_skip_list_sorted_d58

import "testing"

func TestUnit_SkipChunk(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	idx := NewSortedIndex([]int64{15, 3, 9, 21, 9, 12})
	if c := idx.CountRange(9, 15); c != 4 {
		t.Fatalf("expected 4, got %d", c)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
