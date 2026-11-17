package mod_quadtree_spatial_index

import "sync"

func parallelRangeCounts(qt *Quadtree, boxes []Rect) []int {
	out := make([]int, len(boxes))
	var wg sync.WaitGroup
	for i, b := range boxes {
		wg.Add(1)
		go func(idx int, box Rect) {
			defer wg.Done()
			out[idx] = len(qt.QueryRange(box))
		}(i, b)
	}
	wg.Wait()
	return out
}
