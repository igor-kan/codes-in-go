package mod_tarjan_scc_condense_go

import "testing"

func TestProp_TarjanGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	p := TarjanSCC(3, [][2]int{{0, 1}, {1, 2}})
	if p.Count != 3 {
		t.Fatal("DAG of 3 nodes must have 3 SCCs")
	}
}
