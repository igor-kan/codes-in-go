package mod_trie_wildcard_d59

import "testing"

func TestProp_Wtrie(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	wt := NewWildTrie()
	wt.Add("golang")
	if !wt.Search("g.l.ng") || wt.Search("g...") {
		t.Fatalf("length invariant failed")
	}
}
