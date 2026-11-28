package mod_eertree_palindrome_go_d61

import "testing"

func TestProp_Eertree(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	samples := []string{"aabaa", "racecar", "level", "xyzzyx", "abacabadabacaba"}
	for _, s := range samples {
		seen := map[string]bool{}
		for i := 0; i < len(s); i++ {
			for j := i + 1; j <= len(s); j++ {
				sub := s[i:j]
				ok := true
				for l, r := 0, len(sub)-1; l < r; l, r = l+1, r-1 {
					if sub[l] != sub[r] {
						ok = false
						break
					}
				}
				if ok {
					seen[sub] = true
				}
			}
		}
		if got := DistinctPalindromes(s); got != len(seen) {
			t.Fatalf("distinct palindrome mismatch for %q: got %d want %d", s, got, len(seen))
		}
	}
	if DistinctPalindromes("aaaaaaaa") != 8 {
		t.Fatal("unary string must have exactly N distinct palindromes")
	}
}
