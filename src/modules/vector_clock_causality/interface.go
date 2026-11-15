package mod_vector_clock_causality

// NewVectorClock allocates a zeroed vector clock for numNodes processes.
func NewVectorClock(numNodes int) (*VectorClock, error) {
	if numNodes <= 0 {
		return nil, ErrVClockInvalid
	}
	return &VectorClock{Ticks: make([]uint64, numNodes)}, nil
}
