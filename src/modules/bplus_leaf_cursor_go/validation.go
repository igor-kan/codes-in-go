package mod_bplus_leaf_cursor_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrBPlusCursorInvalid
	}
	return nil
}
