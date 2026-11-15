package mod_vector_clock_causality

// CausalRelation classifies partial ordering between two vector clocks.
type CausalRelation int

const (
	RelEqual CausalRelation = iota
	RelBefore
	RelAfter
	RelConcurrent
)

// VectorClock stores per-node monotonically increasing event counters.
type VectorClock struct {
	Ticks []uint64
}
