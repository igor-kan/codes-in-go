package mod_hyperloglog_cardinality

import "sync"

// ParallelHLLMerge shards input streams across goroutines and unions them lock-free.
func ParallelHLLMerge(p uint8, seed uint64, shards [][]uint64) (*HLL, error) {
	locals := make([]*HLL, len(shards))
	var wg sync.WaitGroup
	for i, sh := range shards {
		idx := i
		data := sh
		loc, err := NewHLL(p, seed)
		if err != nil {
			return nil, err
		}
		locals[idx] = loc
		wg.Add(1)
		go func() {
			defer wg.Done()
			AddBatch(locals[idx], data)
		}()
	}
	wg.Wait()
	acc, err := NewHLL(p, seed)
	if err != nil {
		return nil, err
	}
	for _, loc := range locals {
		if err := acc.Merge(loc); err != nil {
			return nil, err
		}
	}
	return acc, nil
}
