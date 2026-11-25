package mod_kmp_prefix_scan_d60

import "testing"

func TestUnit_KmpMatcher(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	hits := KmpFindAll("ababaababa", "aba")
	if len(hits) != 4 || hits[0] != 0 || hits[1] != 2 || hits[2] != 5 || hits[3] != 7 {
		t.Fatalf("unexpected KMP hits: %v", hits)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
