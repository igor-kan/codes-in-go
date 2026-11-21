package mod_radix_compact_trie_d56

import "testing"

func TestUnit_Trie(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	tr := NewPrefixTrie()
	tr.Insert("apple")
	tr.Insert("app")
	tr.Insert("apricot")
	if tr.CountPrefix("ap") != 3 || tr.CountPrefix("app") != 2 || tr.CountPrefix("bat") != 0 {
		t.Fatalf("bad prefix counts")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
