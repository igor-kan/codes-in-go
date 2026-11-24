package mod_sparse_matrix_csr_d59

import "errors"

var (
	ErrCsrInvalid   = errors.New("sparse_matrix_csr_d59: invalid argument")
	ErrCsrBounds    = errors.New("sparse_matrix_csr_d59: index out of bounds")
	ErrCsrEmpty     = errors.New("sparse_matrix_csr_d59: empty input")
	ErrCsrInvariant = errors.New("sparse_matrix_csr_d59: invariant violated")
)
