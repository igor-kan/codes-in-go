package mod_count_min_heavy_go

import "sync"

func parallelIngest(cms *ConservativeCMS, batches [][]string) {
	var wg sync.WaitGroup
	for _, b := range batches {
		wg.Add(1)
		go func(items []string) {
			defer wg.Done()
			for _, it := range items {
				cms.Add(it, 1)
			}
		}(b)
	}
	wg.Wait()
}
