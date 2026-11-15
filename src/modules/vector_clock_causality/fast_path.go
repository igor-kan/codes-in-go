package mod_vector_clock_causality

// HappensBefore returns true iff a strictly precedes b causally.
func HappensBefore(a, b *VectorClock) bool {
	rel, err := a.Compare(b)
	return err == nil && rel == RelBefore
}
