package mod_booth_min_rotation_go_d61

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrBoothRotationInvalid
	}
	return nil
}
