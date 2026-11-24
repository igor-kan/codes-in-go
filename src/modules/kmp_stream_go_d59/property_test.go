package mod_kmp_stream_go_d59

import "testing"

func TestProp_KmpGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	tStr, pStr := "mississippi", "ssi"
	for _, idx := range KmpFindAll(tStr, pStr) {
		if tStr[idx:idx+len(pStr)] != pStr {
			t.Fatalf("false positive hit at %d", idx)
		}
	}
}
