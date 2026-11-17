package mod_aho_corasick_go

import "sync"

func parallelCountCorpus(ac *Automaton, corpus [][]byte) []int {
	out := make([]int, len(corpus))
	var wg sync.WaitGroup
	for i, doc := range corpus {
		wg.Add(1)
		go func(idx int, b []byte) {
			defer wg.Done()
			out[idx] = ac.CountMatches(b)
		}(i, doc)
	}
	wg.Wait()
	return out
}
