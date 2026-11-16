package mod_backoff_jitter_decorrelated

import (
	"sync"
	"time"
)

func ParallelRetryEnvelopes(baseSeed uint64, workers, attempts int) ([][]time.Duration, error) {
	if workers <= 0 || attempts <= 0 {
		return nil, ErrBackoffBounds
	}
	out := make([][]time.Duration, workers)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			cfg := BackoffConfig{
				Base: 10 * time.Millisecond,
				Cap:  5 * time.Second,
				Mode: DecorrelatedJitter,
				Seed: baseSeed + uint64(idx)*101,
			}
			seq, _ := ComputeScheduleSample(cfg, attempts)
			out[idx] = seq
		}(w)
	}
	wg.Wait()
	return out, nil
}
