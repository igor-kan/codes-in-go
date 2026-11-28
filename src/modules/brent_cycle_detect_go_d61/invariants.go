package mod_brent_cycle_detect_go_d61

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrBrentCycleInvariant
	}
	return nil
}
