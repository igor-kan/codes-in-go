package mod_sparse_matrix_csr_d59

type SparseOp interface {
	SpMV(x []float64) []float64
}

type TaggedComponent interface {
	ModuleSlug() string
}
