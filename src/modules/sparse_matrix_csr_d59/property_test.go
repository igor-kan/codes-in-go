package mod_sparse_matrix_csr_d59

import "testing"

func TestProp_Csr(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	id := &CsrMatrix{Rows: 2, Cols: 2, RowPtr: []int{0, 1, 2}, ColIdx: []int{0, 1}, Vals: []float64{1, 1}}
	y := id.SpMV([]float64{7, 9})
	if y[0] != 7 || y[1] != 9 {
		t.Fatalf("identity violated")
	}
}
