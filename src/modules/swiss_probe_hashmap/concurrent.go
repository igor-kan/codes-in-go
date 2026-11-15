package mod_swiss_probe_hashmap

import "sync"

// ShardedSwissMap partitions uint64 keys across RWMutex-guarded SwissMaps.
type ShardedSwissMap struct {
	Shards []*SwissMap
	Locks  []sync.RWMutex
}

// NewShardedSwissMap initializes N shards.
func NewShardedSwissMap(shards int) (*ShardedSwissMap, error) {
	if shards < 1 {
		return nil, ErrSwissInvalid
	}
	s := &ShardedSwissMap{
		Shards: make([]*SwissMap, shards),
		Locks:  make([]sync.RWMutex, shards),
	}
	for i := 0; i < shards; i++ {
		m, _ := NewSwissMap(2)
		s.Shards[i] = m
	}
	return s, nil
}
