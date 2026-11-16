package mod_backoff_jitter_decorrelated

import "time"

func ComputeScheduleSample(cfg BackoffConfig, attempts int) ([]time.Duration, error) {
	s, err := NewBackoffScheduler(cfg)
	if err != nil {
		return nil, err
	}
	out := make([]time.Duration, attempts)
	for i := 0; i < attempts; i++ {
		out[i] = s.Next(i)
	}
	return out, nil
}
