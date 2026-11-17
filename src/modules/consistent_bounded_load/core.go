package mod_consistent_bounded_load

import (
	"math"
	"sort"
	"sync"
)

type BoundedHashRing struct {
	mu       sync.Mutex
	rings    []vnode
	loads    map[string]int
	servers  []string
	factor   float64
	totalReq int
}

func fnv1a64(s string) uint64 {
	var h uint64 = 14695981039346656037
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211
	}
	return h
}

func NewBoundedHashRing(servers []string, replicas int, loadFactor float64) (*BoundedHashRing, error) {
	if err := validateRingCfg(servers, replicas, loadFactor); err != nil {
		return nil, err
	}
	r := &BoundedHashRing{
		loads:   make(map[string]int, len(servers)),
		servers: append([]string(nil), servers...),
		factor:  loadFactor,
	}
	for _, srv := range servers {
		r.loads[srv] = 0
		for rep := 0; rep < replicas; rep++ {
			r.rings = append(r.rings, vnode{hash: fnv1a64(srv + "#" + string(rune('A'+rep))), server: srv})
		}
	}
	sort.Slice(r.rings, func(i, j int) bool { return r.rings[i].hash < r.rings[j].hash })
	return r, nil
}

func (r *BoundedHashRing) currentCap() int {
	avg := float64(r.totalReq+1) / float64(len(r.servers))
	return int(math.Ceil(r.factor * avg))
}

func (r *BoundedHashRing) Acquire(key string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	h := fnv1a64(key)
	start := sort.Search(len(r.rings), func(i int) bool { return r.rings[i].hash >= h }) % len(r.rings)
	capLimit := r.currentCap()
	for step := 0; step < len(r.rings); step++ {
		srv := r.rings[(start+step)%len(r.rings)].server
		if r.loads[srv] < capLimit {
			r.loads[srv]++
			r.totalReq++
			return srv, nil
		}
	}
	return "", ErrBoundedRingBounds
}

func (r *BoundedHashRing) Release(server string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	c, ok := r.loads[server]
	if !ok || c == 0 {
		return ErrBoundedRingInvalid
	}
	r.loads[server] = c - 1
	r.totalReq--
	return nil
}

func (r *BoundedHashRing) LoadSnapshot() (map[string]int, int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := make(map[string]int, len(r.loads))
	for k, v := range r.loads {
		cp[k] = v
	}
	return cp, r.totalReq
}
