package mod_singleflight_dedup_group

func VerifySingleExecution(execs int64, values []int64) error {
	if execs != 1 || len(values) == 0 {
		return ErrFlightInvariant
	}
	for _, v := range values {
		if v != values[0] {
			return ErrFlightInvariant
		}
	}
	return nil
}
