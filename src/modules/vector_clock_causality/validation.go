package mod_vector_clock_causality

// ValidateVectorClock checks clock non-emptiness.
func ValidateVectorClock(vc *VectorClock) error {
	if vc == nil || len(vc.Ticks) == 0 {
		return ErrVClockEmpty
	}
	return nil
}
