package mod_knuth_morris_pratt_go

import "testing"

func TestProp_KmpGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	if len(KmpSearchAll("abcdef", "xyz")) != 0 { t.Fatal("expected zero hits") }
}
