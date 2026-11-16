package mod_backoff_jitter_decorrelated

import "time"

func VerifyDelayBounds(cfg BackoffConfig, delays []time.Duration) error {
	for _, d := range delays {
		if d < 0 || d > cfg.Cap {
			return ErrBackoffInvariant
		}
		if cfg.Mode == DecorrelatedJitter && d < cfg.Base {
			return ErrBackoffInvariant
		}
	}
	return nil
}
