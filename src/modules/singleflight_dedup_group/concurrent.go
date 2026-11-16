package mod_singleflight_dedup_group

import (
	"sync"
	"sync/atomic"
	"time"
)

func RunStampedBurst(key string, callers int) (executions int64, values []int64) {
	g := NewDedupGroup()
	values = make([]int64, callers)
	var execs atomic.Int64
	var ready sync.WaitGroup
	ready.Add(callers)
	start := make(chan struct{})
	var done sync.WaitGroup
	for i := 0; i < callers; i++ {
		done.Add(1)
		go func(idx int) {
			defer done.Done()
			ready.Done()
			<-start
			res := g.Do(key, func() (int64, error) {
				n := execs.Add(1)
				time.Sleep(15 * time.Millisecond)
				return 4200 + n, nil
			})
			values[idx] = res.Val
		}(i)
	}
	ready.Wait()
	close(start)
	done.Wait()
	return execs.Load(), values
}
