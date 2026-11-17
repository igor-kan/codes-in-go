package mod_consistent_bounded_load

import (
	"strconv"
	"sync"
)

func parallelAcquireBatch(ring *BoundedHashRing, keys []string) []string {
	out := make([]string, len(keys))
	var wg sync.WaitGroup
	for i, k := range keys {
		wg.Add(1)
		go func(idx int, key string) {
			defer wg.Done()
			srv, _ := ring.Acquire(key + strconv.Itoa(idx))
			out[idx] = srv
		}(i, k)
	}
	wg.Wait()
	return out
}
