package mod_aho_multi_scan_d59

import "testing"

func TestUnit_AcGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	ac := BuildAc([]string{"he", "she", "hers", "his"})
	if ac.CountHits("ahishers") != 4 {
		t.Fatalf("bad ac hits")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
