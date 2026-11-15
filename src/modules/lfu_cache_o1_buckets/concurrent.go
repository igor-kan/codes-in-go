package mod_lfu_cache_o1_buckets

import "sync"

// SyncLFU wraps LFUCache under a mutex for concurrent goroutine access.
type SyncLFU struct {
	mu    sync.Mutex
	cache *LFUCache
}

// PutSafe acquires mutex and delegates Put.
func (s *SyncLFU) PutSafe(k, v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache.Put(k, v)
}
