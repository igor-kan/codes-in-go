package mod_interval_coalesce_go

import "sync"

func parallelProbePoints(ledger *DisjointIntervals, pts []int64) []bool {
	out := make([]bool, len(pts))
	var wg sync.WaitGroup
	for i, p := range pts {
		wg.Add(1)
		go func(idx int, val int64) {
			defer wg.Done()
			out[idx] = ledger.Contains(val)
		}(i, p)
	}
	wg.Wait()
	return out
}
