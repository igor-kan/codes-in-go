package mod_suffix_array_kasai_d60

import "testing"

func TestUnit_SuffixIndex(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	sa, lcp := BuildSuffixAndLCP("banana")
	if len(sa) != 6 || lcp[2] != 3 {
		t.Fatalf("unexpected sa/lcp for banana: %v %v", sa, lcp)
	}
	if DistinctSubstrings("abab") != 7 {
		t.Fatalf("expected 7 distinct substrings for abab, got %d", DistinctSubstrings("abab"))
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
