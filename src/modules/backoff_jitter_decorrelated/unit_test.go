package mod_backoff_jitter_decorrelated

import (
	"testing"
	"time"
)

func TestBackoffBoundsAndReset(t *testing.T) {
	cfg := BackoffConfig{Base: 10 * time.Millisecond, Cap: 1 * time.Second, Mode: FullJitter, Seed: 7}
	s, err := NewBackoffScheduler(cfg)
	if err != nil {
		t.Fatal(err)
	}
	var seq []time.Duration
	for i := 0; i < 20; i++ {
		seq = append(seq, s.Next(i))
	}
	if err := VerifyDelayBounds(cfg, seq); err != nil {
		t.Fatal(err)
	}
}
