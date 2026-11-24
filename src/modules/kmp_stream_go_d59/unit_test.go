package mod_kmp_stream_go_d59

import "testing"

func TestUnit_KmpGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	hits := KmpFindAll("abababa", "aba")
	if len(hits) != 3 || hits[0] != 0 || hits[1] != 2 || hits[2] != 4 {
		t.Fatalf("bad kmp hits: %v", hits)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
