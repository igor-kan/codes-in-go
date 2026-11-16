package mod_backoff_jitter_decorrelated

import (
	"math/rand/v2"
	"time"
)

type BackoffScheduler struct {
	cfg  BackoffConfig
	rng  *rand.Rand
	prev time.Duration
}

func NewBackoffScheduler(cfg BackoffConfig) (*BackoffScheduler, error) {
	if err := ValidateBackoffConfig(cfg); err != nil {
		return nil, err
	}
	src := rand.NewPCG(cfg.Seed, cfg.Seed^0x9e3779b97f4a7c15)
	return &BackoffScheduler{
		cfg:  cfg,
		rng:  rand.New(src),
		prev: cfg.Base,
	}, nil
}

func (b *BackoffScheduler) Reset() {
	b.prev = b.cfg.Base
}

func (b *BackoffScheduler) Next(attempt int) time.Duration {
	if attempt < 0 {
		attempt = 0
	}
	if b.cfg.Mode == DecorrelatedJitter {
		upper := b.prev * 3
		if upper > b.cfg.Cap {
			upper = b.cfg.Cap
		}
		span := upper - b.cfg.Base
		var sleep time.Duration
		if span <= 0 {
			sleep = b.cfg.Base
		} else {
			sleep = b.cfg.Base + time.Duration(b.rng.Int64N(int64(span)+1))
		}
		if sleep > b.cfg.Cap {
			sleep = b.cfg.Cap
		}
		b.prev = sleep
		return sleep
	}
	ceil := b.cfg.Base
	for i := 0; i < attempt && ceil < b.cfg.Cap; i++ {
		if ceil > b.cfg.Cap/2 {
			ceil = b.cfg.Cap
			break
		}
		ceil *= 2
	}
	if ceil <= 0 {
		return 0
	}
	return time.Duration(b.rng.Int64N(int64(ceil) + 1))
}
