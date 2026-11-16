package mod_crdt_lww_element_set

// ReplicatedSet defines element mutation, state inspection, and semilattice merge.
type ReplicatedSet interface {
	Add(elem string, ts uint64) error
	Remove(elem string, ts uint64) error
	Contains(elem string) bool
	Merge(other *LWWSet) error
}

// NewLWWSet constructs an empty LWW-Element-Set with the given tie-break bias.
func NewLWWSet(bias BiasMode) *LWWSet {
	return &LWWSet{
		Bias:    bias,
		Adds:    make(map[string]uint64),
		Removes: make(map[string]uint64),
	}
}
