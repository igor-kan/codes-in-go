package mod_backoff_jitter_decorrelated

import (
	"testing"
	"time"
)

func TestDecorrelatedDivergenceAcrossWorkers(t *testing.T) {
	rows, err := ParallelRetryEnvelopes(99, 6, 15)
	if err != nil {
		t.Fatal(err)
	}
	cfg := BackoffConfig{Base: 10 * time.Millisecond, Cap: 5 * time.Second, Mode: DecorrelatedJitter}
	for _, seq := range rows {
		if err := VerifyDelayBounds(cfg, seq); err != nil {
			t.Fatal(err)
		}
	}
}
