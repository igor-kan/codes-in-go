package mod_rw_fair_rwlock_queue

import (
	"sync"
	"sync/atomic"
)

func StressFairRW(readers, writers, ops int) (int64, error) {
	m := NewFairRWMutex()
	shared := int64(0)
	var activeWriters atomic.Int32
	var activeReads atomic.Int32
	var raceHit atomic.Bool
	var wg sync.WaitGroup
	for w := 0; w < writers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < ops; i++ {
				m.Lock()
				if activeWriters.Add(1) != 1 || activeReads.Load() != 0 {
					raceHit.Store(true)
				}
				shared++
				activeWriters.Add(-1)
				m.Unlock()
			}
		}()
	}
	for r := 0; r < readers; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < ops; i++ {
				m.RLock()
				activeReads.Add(1)
				if activeWriters.Load() != 0 || shared < 0 {
					raceHit.Store(true)
				}
				activeReads.Add(-1)
				m.RUnlock()
			}
		}()
	}
	wg.Wait()
	if raceHit.Load() || shared != int64(writers*ops) {
		return shared, ErrFairRWInvariant
	}
	return shared, nil
}
