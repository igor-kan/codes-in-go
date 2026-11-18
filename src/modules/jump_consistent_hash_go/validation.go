package mod_jump_consistent_hash_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrJumpHashInvalid
	}
	return nil
}
