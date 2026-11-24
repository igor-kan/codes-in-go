package mod_dijkstra_sp_go_d59

import "sync"

func parallelFoldSlices(a, b []int64) int64 {
	var wg sync.WaitGroup
	var ra, rb int64
	wg.Add(2)
	go func() { defer wg.Done(); ra = fastFoldSlice(a) }()
	go func() { defer wg.Done(); rb = fastFoldSlice(b) }()
	wg.Wait()
	return ra ^ rb
}
