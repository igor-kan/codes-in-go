package mod_union_find_weighted

import "sync"

// GuardedDSU synchronizes concurrent constraint checks across goroutines.
type GuardedDSU struct {
	mu  sync.Mutex
	dsu *WeightedDSU
}

// RelateSync applies a relation safely under mutex.
func (g *GuardedDSU) RelateSync(x, y int, diff float64) (bool, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.dsu.Relate(x, y, diff)
}
