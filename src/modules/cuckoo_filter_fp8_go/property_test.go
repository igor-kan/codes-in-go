package mod_cuckoo_filter_fp8_go

import "testing"

func TestProp_CuckooGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	cf := NewCuckooFilter(16)
	cf.Insert(999)
	if !cf.Lookup(999) {
		t.Fatal("missing 999")
	}
}
