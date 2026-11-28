package mod_myers_bit_edit_go_d61

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrMyersEditInvariant
	}
	return nil
}
