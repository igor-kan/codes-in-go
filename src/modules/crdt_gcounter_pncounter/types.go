package mod_crdt_gcounter_pncounter

// GCounter tracks monotonic per-replica counts via pointwise maximum semilattice.
type GCounter struct {
	Slots []uint64
}

// PNCounter pairs increment (P) and decrement (N) GCounters for bidirectional counts.
type PNCounter struct {
	P GCounter
	N GCounter
}
