package mod_consistent_ring_go_d60

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrHashRingInvalid
	}
	return nil
}
