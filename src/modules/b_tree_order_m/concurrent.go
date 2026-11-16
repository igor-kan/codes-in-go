package mod_b_tree_order_m

import "sync"

func ParallelBTreeLookup(t *BTree, keys []int64) []bool {
	found := make([]bool, len(keys))
	if len(keys) == 0 {
		return found
	}
	var wg sync.WaitGroup
	workers := 4
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := id; i < len(keys); i += workers {
				_, found[i] = t.Get(keys[i])
			}
		}(w)
	}
	wg.Wait()
	return found
}
