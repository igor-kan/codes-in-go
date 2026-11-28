package mod_boruvka_mst_go_d61

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrBoruvkaMstInvalid
	}
	return nil
}
