package mod_sparse_matrix_csr_d59

import "testing"

func TestUnit_Csr(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	m := &CsrMatrix{Rows: 2, Cols: 3, RowPtr: []int{0, 2, 3}, ColIdx: []int{0, 2, 1}, Vals: []float64{2, 5, 3}}
	y := m.SpMV([]float64{10, 4, 1})
	if y[0] != 25 || y[1] != 12 {
		t.Fatalf("bad spmv: %v", y)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
