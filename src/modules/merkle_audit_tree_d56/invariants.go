package mod_merkle_audit_tree_d56

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrMerkleInvariant
	}
	return nil
}
