package mod_sparse_matrix_csr_d59

type CsrMatrix struct {
	Rows   int
	Cols   int
	RowPtr []int
	ColIdx []int
	Vals   []float64
}

type ModuleMeta struct {
	Slug    string
	Version int
}
