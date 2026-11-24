package mod_sparse_matrix_csr_d59

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrCsrInvalid
	}
	return nil
}
