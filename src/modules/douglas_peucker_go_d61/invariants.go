package mod_douglas_peucker_go_d61

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrPolylineRdpInvariant
	}
	return nil
}
