package mod_aho_corasick_go

import "testing"

func TestAhoCorasickMatches(t *testing.T) {
	pats := [][]byte{[]byte("he"), []byte("she"), []byte("his"), []byte("hers")}
	ac, err := CompilePatterns(pats)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyAcMatchesNaive(ac, pats, []byte("ahishers_she")); err != nil {
		t.Fatal(err)
	}
}
