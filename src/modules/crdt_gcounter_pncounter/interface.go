package mod_crdt_gcounter_pncounter

// ReplicatedCounter defines replica increment/decrement, query, and lattice merge.
type ReplicatedCounter interface {
	Value() int64
	Merge(other *PNCounter) error
}

// NewPNCounter allocates a PN-Counter across numReplicas replicas.
func NewPNCounter(numReplicas int) (*PNCounter, error) {
	if numReplicas <= 0 {
		return nil, ErrCRDTInvalid
	}
	return &PNCounter{
		P: GCounter{Slots: make([]uint64, numReplicas)},
		N: GCounter{Slots: make([]uint64, numReplicas)},
	}, nil
}
