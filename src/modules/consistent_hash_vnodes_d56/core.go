package mod_consistent_hash_vnodes_d56

import (
	"sort"
	"strconv"
)

func fnv32a(s string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

type HashRing struct {
	replicas int
	ring     []VNode
}

func NewHashRing(hosts []string, replicas int) *HashRing {
	r := &HashRing{replicas: replicas}
	for _, h := range hosts {
		for v := 0; v < replicas; v++ {
			tok := fnv32a(h + "#" + strconv.Itoa(v))
			r.ring = append(r.ring, VNode{Token: tok, Host: h})
		}
	}
	sort.Slice(r.ring, func(i, j int) bool { return r.ring[i].Token < r.ring[j].Token })
	return r
}

func (r *HashRing) Lookup(key string) string {
	if len(r.ring) == 0 {
		return ""
	}
	tok := fnv32a(key)
	idx := sort.Search(len(r.ring), func(i int) bool { return r.ring[i].Token >= tok })
	if idx == len(r.ring) {
		idx = 0
	}
	return r.ring[idx].Host
}
