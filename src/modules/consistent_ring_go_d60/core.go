package mod_consistent_ring_go_d60

import (
	"hash/fnv"
	"sort"
	"strconv"
)

type ConsistentRing struct {
	replicas int
	slots    []RingSlot
	active   map[string]bool
}

func NewConsistentRing(replicas int) *ConsistentRing {
	if replicas <= 0 {
		replicas = 16
	}
	return &ConsistentRing{replicas: replicas, active: make(map[string]bool)}
}

func ringHash64(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

func (r *ConsistentRing) AddNode(node string) {
	if r.active[node] {
		return
	}
	r.active[node] = true
	for i := 0; i < r.replicas; i++ {
		tok := ringHash64(node + "#" + strconv.Itoa(i))
		r.slots = append(r.slots, RingSlot{Token: tok, Node: node})
	}
	sort.Slice(r.slots, func(i, j int) bool { return r.slots[i].Token < r.slots[j].Token })
}

func (r *ConsistentRing) RemoveNode(node string) {
	delete(r.active, node)
	kept := r.slots[:0]
	for _, sl := range r.slots {
		if sl.Node != node {
			kept = append(kept, sl)
		}
	}
	r.slots = kept
}

func (r *ConsistentRing) RouteKey(key string) (string, bool) {
	if len(r.slots) == 0 {
		return "", false
	}
	tok := ringHash64(key)
	idx := sort.Search(len(r.slots), func(i int) bool { return r.slots[i].Token >= tok })
	if idx == len(r.slots) {
		idx = 0
	}
	return r.slots[idx].Node, true
}
