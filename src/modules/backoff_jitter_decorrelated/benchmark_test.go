package mod_backoff_jitter_decorrelated

import (
	"testing"
	"time"
)

func BenchmarkDecorrelatedNext(b *testing.B) {
	s, _ := NewBackoffScheduler(BackoffConfig{Base: 5 * time.Millisecond, Cap: time.Second, Mode: DecorrelatedJitter, Seed: 42})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Next(i & 15)
	}
}
