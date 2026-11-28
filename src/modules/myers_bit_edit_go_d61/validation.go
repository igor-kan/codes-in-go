package mod_myers_bit_edit_go_d61

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrMyersEditInvalid
	}
	return nil
}
