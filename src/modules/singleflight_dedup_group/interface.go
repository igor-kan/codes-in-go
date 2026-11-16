package mod_singleflight_dedup_group

// Coalescer suppresses duplicate concurrent executions per logical request key.
type Coalescer interface {
	Do(key string, fn func() (int64, error)) GroupResult
	Forget(key string)
}
