package mod_backoff_jitter_decorrelated

func ValidateBackoffConfig(cfg BackoffConfig) error {
	if cfg.Base <= 0 || cfg.Cap < cfg.Base {
		return ErrBackoffInvalid
	}
	return nil
}
