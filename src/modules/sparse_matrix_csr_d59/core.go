package mod_sparse_matrix_csr_d59

func (m *CsrMatrix) SpMV(x []float64) []float64 {
	y := make([]float64, m.Rows)
	for r := 0; r < m.Rows; r++ {
		var s float64
		for idx := m.RowPtr[r]; idx < m.RowPtr[r+1]; idx++ {
			s += m.Vals[idx] * x[m.ColIdx[idx]]
		}
		y[r] = s
	}
	return y
}
