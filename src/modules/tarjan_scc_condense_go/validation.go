package mod_tarjan_scc_condense_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrTarjanGoInvalid
	}
	return nil
}
