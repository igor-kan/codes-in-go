package mod_crdt_lww_element_set

import "sync"

func ConvergeReplicas(replicas []*LWWSet) (*LWWSet, error) {
	if len(replicas) == 0 {
		return nil, ErrLWWEmpty
	}
	var mu sync.Mutex
	merged := NewLWWSet(replicas[0].Bias)
	var wg sync.WaitGroup
	var firstErr error
	for _, r := range replicas {
		wg.Add(1)
		go func(rep *LWWSet) {
			defer wg.Done()
			cl := rep.Clone()
			mu.Lock()
			if err := merged.Merge(cl); err != nil && firstErr == nil {
				firstErr = err
			}
			mu.Unlock()
		}(r)
	}
	wg.Wait()
	return merged, firstErr
}
