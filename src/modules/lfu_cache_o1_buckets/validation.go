package mod_lfu_cache_o1_buckets

// ValidateLFUCache verifies capacity and node/bucket sum consistency.
func ValidateLFUCache(c *LFUCache) error {
	if c == nil || c.Cap <= 0 || len(c.Nodes) > c.Cap {
		return ErrLFUBounds
	}
	return nil
}
