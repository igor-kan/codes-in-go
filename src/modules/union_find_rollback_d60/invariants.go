package mod_union_find_rollback_d60

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrRollbackDsuInvariant
	}
	return nil
}
