package mod_merkle_audit_tree_d56

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrMerkleInvalid
	}
	return nil
}
