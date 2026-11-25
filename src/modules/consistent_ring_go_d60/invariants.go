package mod_consistent_ring_go_d60

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrHashRingInvariant
	}
	return nil
}
