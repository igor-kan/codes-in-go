package mod_hyperloglog_cardinality

// CardinalitySketch defines streaming observation, estimation, and semilattice merge.
type CardinalitySketch interface {
	Add(item uint64)
	Count() float64
	Merge(other *HLL) error
}

// NewHLL constructs an HLL sketch with 4 <= p <= 16.
func NewHLL(p uint8, seed uint64) (*HLL, error) {
	if p < 4 || p > 16 {
		return nil, ErrHLLInvalid
	}
	return &HLL{
		Precision: p,
		Seed:      seed,
		Registers: make([]uint8, 1<<p),
	}, nil
}
