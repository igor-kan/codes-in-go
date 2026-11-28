package mod_boruvka_mst_go_d61

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrBoruvkaMstInvariant
	}
	return nil
}
