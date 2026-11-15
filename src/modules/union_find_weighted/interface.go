package mod_union_find_weighted

// NewWeightedDSU allocates n singleton potential sets.
func NewWeightedDSU(n int) (*WeightedDSU, error) {
	if n <= 0 {
		return nil, ErrWDSUInvalid
	}
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &WeightedDSU{
		Parent: p,
		Rank:   make([]int, n),
		Weight: make([]float64, n),
	}, nil
}
