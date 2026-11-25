package mod_suffix_array_kasai_d60

import "testing"

func TestProp_SuffixIndex(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	if DistinctSubstrings("aaaa") != 4 {
		t.Fatal("unary string must have N distinct substrings")
	}
	if DistinctSubstrings("abcd") != 10 {
		t.Fatal("all-unique string must have N*(N+1)/2 distinct substrings")
	}
}
