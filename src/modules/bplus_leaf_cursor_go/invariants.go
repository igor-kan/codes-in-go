package mod_bplus_leaf_cursor_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrBPlusCursorInvariant
	}
	return nil
}
