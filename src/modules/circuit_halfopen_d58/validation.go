package mod_circuit_halfopen_d58

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrBreakerInvalid
	}
	return nil
}
