package mod_aho_corasick_go

import "bytes"

func verifyAcMatchesNaive(ac *Automaton, patterns [][]byte, text []byte) error {
	exp := 0
	for _, p := range patterns {
		for off := 0; off+len(p) <= len(text); off++ {
			if bytes.Equal(text[off:off+len(p)], p) {
				exp++
			}
		}
	}
	if ac.CountMatches(text) != exp {
		return ErrAhoCorasickInvariant
	}
	return nil
}
