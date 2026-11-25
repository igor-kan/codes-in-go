package mod_union_find_rollback_d60

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrRollbackDsuInvalid
	}
	return nil
}
