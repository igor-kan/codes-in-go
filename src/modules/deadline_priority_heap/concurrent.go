package mod_deadline_priority_heap

import (
	"strconv"
	"sync"
)

func parallelScheduleBatch(h *DeadlineHeap, count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			_ = h.Schedule("t-"+strconv.Itoa(idx), int64(idx*10))
		}(i)
	}
	wg.Wait()
}
