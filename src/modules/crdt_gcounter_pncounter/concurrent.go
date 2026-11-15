package mod_crdt_gcounter_pncounter

import "sync"

// ReplicatedUpdate applies replica-local concurrent increments cleanly with final merge.
func ReplicatedUpdate(replicas int, ops [][2]int64) (*PNCounter, error) {
	nodes := make([]*PNCounter, replicas)
	for i := 0; i < replicas; i++ {
		nodes[i], _ = NewPNCounter(replicas)
	}
	var wg sync.WaitGroup
	for r := 0; r < replicas; r++ {
		rid := r
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx, op := range ops {
				if idx%replicas == rid {
					_ = nodes[rid].Add(rid, op[1])
				}
			}
		}()
	}
	wg.Wait()
	merged, _ := NewPNCounter(replicas)
	for _, n := range nodes {
		_ = merged.Merge(n)
	}
	return merged, nil
}
