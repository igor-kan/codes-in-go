package mod_lfu_cache_o1_buckets

// VerifyLFUInvariants checks bijection between Nodes map and Buckets contents.
func VerifyLFUInvariants(c *LFUCache) error {
	if err := ValidateLFUCache(c); err != nil {
		return err
	}
	sum := 0
	minSeen := int(^uint(0) >> 1)
	for freq, b := range c.Buckets {
		if b.Len() == 0 {
			return ErrLFUInvariant
		}
		sum += b.Len()
		if freq < minSeen {
			minSeen = freq
		}
	}
	if sum != len(c.Nodes) {
		return ErrLFUInvariant
	}
	if len(c.Nodes) > 0 && c.MinFreq != minSeen {
		return ErrLFUInvariant
	}
	return nil
}
