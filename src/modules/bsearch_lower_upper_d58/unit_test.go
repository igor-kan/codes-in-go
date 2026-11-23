package mod_bsearch_lower_upper_d58

import "testing"

func TestUnit_BSearch(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	arr := []int64{1, 3, 3, 3, 7, 9}
	if CountSortedOccurrences(arr, 3) != 3 || CountSortedOccurrences(arr, 5) != 0 {
		t.Fatalf("bad occurrence count")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
