package mod_myers_bit_edit_go_d61

import "testing"

func TestProp_MyersEdit(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	pairs := [][2]string{{"kitten", "sitting"}, {"flaw", "lawn"}, {"intention", "execution"}, {"aaaa", "aa"}, {"abc", ""}}
	for _, pr := range pairs {
		if MyersDistance(pr[0], pr[1]) != classicDistance(pr[0], pr[1]) {
			t.Fatalf("bit edit mismatch for %q vs %q", pr[0], pr[1])
		}
		if MyersDistance(pr[0], pr[1]) != MyersDistance(pr[1], pr[0]) {
			t.Fatalf("Levenshtein symmetry violated for %q vs %q", pr[0], pr[1])
		}
	}
}
