package mod_circuit_halfopen_d58

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrBreakerInvariant
	}
	return nil
}
