package mod_crdt_lww_element_set

// BiasMode selects deterministic tie-breaking when add and remove timestamps match.
type BiasMode uint8

const (
	BiasRemoveWins BiasMode = iota
	BiasAddWins
)

// LWWSet stores monotonic add and remove tombstone timestamps per key.
type LWWSet struct {
	Bias    BiasMode
	Adds    map[string]uint64
	Removes map[string]uint64
}
