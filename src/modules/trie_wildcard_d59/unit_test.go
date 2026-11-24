package mod_trie_wildcard_d59

import "testing"

func TestUnit_Wtrie(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	wt := NewWildTrie()
	wt.Add("bad"); wt.Add("dad"); wt.Add("mad")
	if !wt.Search(".ad") || !wt.Search("b..") || wt.Search("pad") {
		t.Fatalf("wildcard search failed")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
