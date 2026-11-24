package mod_aho_multi_scan_d59

import "testing"

func TestProp_AcGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	ac := BuildAc([]string{"aa", "aaa"})
	if ac.CountHits("aaaa") != 5 {
		t.Fatalf("overlap count failed")
	}
}
