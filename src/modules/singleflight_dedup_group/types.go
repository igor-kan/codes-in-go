package mod_singleflight_dedup_group

// GroupResult conveys leader execution output and whether the result was shared across callers.
type GroupResult struct {
	Val    int64
	Err    error
	Shared bool
}
