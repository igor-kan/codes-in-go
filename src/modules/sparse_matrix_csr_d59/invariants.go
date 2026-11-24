package mod_sparse_matrix_csr_d59

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrCsrInvariant
	}
	return nil
}
