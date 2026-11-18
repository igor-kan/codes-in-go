package mod_tarjan_scc_condense_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrTarjanGoInvariant
	}
	return nil
}
