package mod_lfu_cache_o1_buckets

import "container/list"

// NewLFUCache creates an LFU cache bounded to positive capacity cap.
func NewLFUCache(cap int) (*LFUCache, error) {
	if cap <= 0 {
		return nil, ErrLFUInvalid
	}
	return &LFUCache{
		Cap:     cap,
		Nodes:   make(map[int]*list.Element, cap),
		Buckets: make(map[int]*list.List),
	}, nil
}
