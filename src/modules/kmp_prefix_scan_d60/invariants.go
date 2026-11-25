package mod_kmp_prefix_scan_d60

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrKmpMatcherInvariant
	}
	return nil
}
