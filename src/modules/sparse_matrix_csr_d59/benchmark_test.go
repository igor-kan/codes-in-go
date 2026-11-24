package mod_sparse_matrix_csr_d59

import "testing"

func Benchmark_Csr(b *testing.B) {
	matB := &CsrMatrix{Rows: 2, Cols: 2, RowPtr: []int{0, 1, 2}, ColIdx: []int{0, 1}, Vals: []float64{2, 3}}
	xVec := []float64{4, 5}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = matB.SpMV(xVec)
	}
	_ = sink
}
