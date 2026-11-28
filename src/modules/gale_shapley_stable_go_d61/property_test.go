package mod_gale_shapley_stable_go_d61

import "testing"

func TestProp_StableMarriage(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	pa := [][]int{{0, 1, 2, 3}, {1, 0, 2, 3}, {2, 1, 0, 3}, {0, 2, 1, 3}}
	pb := [][]int{{3, 1, 2, 0}, {1, 3, 0, 2}, {2, 0, 3, 1}, {0, 1, 2, 3}}
	m := StableMatch(pa, pb)
	seen := map[int]bool{}
	matchB := make([]int, 4)
	for a := 0; a < 4; a++ {
		b := m[a]
		if b < 0 || b >= 4 || seen[b] {
			t.Fatalf("not a perfect matching: %v", m)
		}
		seen[b] = true
		matchB[b] = a
	}
	rankB := make([][]int, 4)
	for b := 0; b < 4; b++ {
		rankB[b] = make([]int, 4)
		for r, a := range pb[b] {
			rankB[b][a] = r
		}
	}
	for a := 0; a < 4; a++ {
		for _, b := range pa[a] {
			if b == m[a] {
				break
			}
			if rankB[b][a] < rankB[b][matchB[b]] {
				t.Fatalf("blocking pair (%d,%d) in %v", a, b, m)
			}
		}
	}
}
