package mod_knuth_morris_pratt_go

import "testing"

func TestUnit_KmpGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	hits := KmpSearchAll("ababaababa", "aba")
	if len(hits) != 4 || hits[0] != 0 || hits[3] != 7 {
		t.Fatalf("unexpected KMP hits: %v", hits)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
