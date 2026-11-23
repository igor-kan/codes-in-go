package mod_worker_fanout_pool_d58

import "sync"

func ParallelSquareMap(xs []int64, workers int) []int64 {
	if workers <= 0 {
		workers = 1
	}
	out := make([]int64, len(xs))
	ch := make(chan int, len(xs))
	for i := range xs {
		ch <- i
	}
	close(ch)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range ch {
				v := xs[idx]
				out[idx] = v * v
			}
		}()
	}
	wg.Wait()
	return out
}
