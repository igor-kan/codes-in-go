package mod_vector_clock_causality

// VerifyTransitivity verifies A->B && B->C => A->C.
func VerifyTransitivity(a, b, c *VectorClock) error {
	if HappensBefore(a, b) && HappensBefore(b, c) && !HappensBefore(a, c) {
		return ErrVClockInvariant
	}
	return nil
}
