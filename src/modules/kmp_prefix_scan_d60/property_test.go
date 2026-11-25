package mod_kmp_prefix_scan_d60

import "testing"

func TestProp_KmpMatcher(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	txt := "xyzabababa"
	pat := "aba"
	for _, idx := range KmpFindAll(txt, pat) {
		if txt[idx:idx+len(pat)] != pat {
			t.Fatalf("invalid match at %d", idx)
		}
	}
}
