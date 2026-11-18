package mod_sliding_percentile_p2_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrWindowQuantInvariant
	}
	return nil
}
