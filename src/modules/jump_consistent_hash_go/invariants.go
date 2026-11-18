package mod_jump_consistent_hash_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrJumpHashInvariant
	}
	return nil
}
