package mod_aho_multi_scan_d59

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrAcGoInvariant
	}
	return nil
}
