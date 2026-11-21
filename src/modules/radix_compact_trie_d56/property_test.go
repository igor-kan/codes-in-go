package mod_radix_compact_trie_d56

import "testing"

func TestProp_Trie(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	tr := NewPrefixTrie()
	tr.Insert("alpha")
	tr.Insert("beta")
	if tr.CountPrefix("") != 2 {
		t.Fatalf("root count mismatch")
	}
}
