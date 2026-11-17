package mod_count_min_heavy_go

import (
	"sort"
	"sync"
)

type ConservativeCMS struct {
	mu      sync.RWMutex
	depth   int
	width   int
	table   [][]uint64
	heavy   map[string]uint64
}

func NewConservativeCMS(depth, width int) (*ConservativeCMS, error) {
	if err := validateSketchDims(depth, width); err != nil {
		return nil, err
	}
	t := make([][]uint64, depth)
	for d := 0; d < depth; d++ {
		t[d] = make([]uint64, width)
	}
	return &ConservativeCMS{depth: depth, width: width, table: t, heavy: make(map[string]uint64)}, nil
}

func hashRow(key string, seed uint64) uint64 {
	h := 14695981039346656037 ^ seed*0x9e3779b97f4a7c15
	for i := 0; i < len(key); i++ {
		h ^= uint64(key[i])
		h *= 1099511628211
	}
	return h
}

func (c *ConservativeCMS) Add(key string, delta uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	idxs := make([]int, c.depth)
	curMin := ^uint64(0)
	for d := 0; d < c.depth; d++ {
		idx := int(hashRow(key, uint64(d+1)) % uint64(c.width))
		idxs[d] = idx
		if c.table[d][idx] < curMin {
			curMin = c.table[d][idx]
		}
	}
	target := curMin + delta
	for d := 0; d < c.depth; d++ {
		if c.table[d][idxs[d]] < target {
			c.table[d][idxs[d]] = target
		}
	}
	c.heavy[key] = target
}

func (c *ConservativeCMS) Estimate(key string) uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	curMin := ^uint64(0)
	for d := 0; d < c.depth; d++ {
		idx := int(hashRow(key, uint64(d+1)) % uint64(c.width))
		if c.table[d][idx] < curMin {
			curMin = c.table[d][idx]
		}
	}
	return curMin
}

func (c *ConservativeCMS) TopK(k int) []HeavyItem {
	c.mu.RLock()
	defer c.mu.RUnlock()
	items := make([]HeavyItem, 0, len(c.heavy))
	for key := range c.heavy {
		items = append(items, HeavyItem{Key: key, Score: c.Estimate(key)})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Score != items[j].Score {
			return items[i].Score > items[j].Score
		}
		return items[i].Key < items[j].Key
	})
	if k < len(items) {
		return items[:k]
	}
	return items
}
