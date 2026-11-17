package mod_aho_corasick_go

import "testing"

func TestAhoCorasickParallelCorpus(t *testing.T) {
	pats := [][]byte{[]byte("ab"), []byte("aba"), []byte("ba")}
	ac, _ := CompilePatterns(pats)
	docs := [][]byte{[]byte("ababa"), []byte("baba"), []byte("zzzz")}
	cnts := parallelCountCorpus(ac, docs)
	if cnts[0] != 6 || cnts[1] != 4 || cnts[2] != 0 {
		t.Fatalf("unexpected match vector: %v", cnts)
	}
}
