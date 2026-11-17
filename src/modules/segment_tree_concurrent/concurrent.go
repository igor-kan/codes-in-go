package mod_segment_tree_concurrent

import "sync"

func parallelRangeQueries(st *ConcurrentSegTree, ranges [][2]int) []SegStats {
	out := make([]SegStats, len(ranges))
	var wg sync.WaitGroup
	for i, rg := range ranges {
		wg.Add(1)
		go func(idx int, l, r int) {
			defer wg.Done()
			out[idx], _ = st.Query(l, r)
		}(i, rg[0], rg[1])
	}
	wg.Wait()
	return out
}
