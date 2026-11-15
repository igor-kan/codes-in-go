package mod_lfu_cache_o1_buckets

// PeekMinFreq returns current minimum frequency or 0 when empty.
func (c *LFUCache) PeekMinFreq() int {
	if len(c.Nodes) == 0 {
		return 0
	}
	return c.MinFreq
}
