package mod_booth_min_rotation_go_d61

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrBoothRotationInvariant
	}
	return nil
}
