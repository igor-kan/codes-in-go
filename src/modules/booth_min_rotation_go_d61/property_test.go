package mod_booth_min_rotation_go_d61

import "testing"

func TestProp_BoothRotation(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	samples := []string{"banana", "cba", "aaaa", "abab", "z", "mississippi", "abcabcabc"}
	for _, s := range samples {
		best := MinimalRotation(s)
		runes := []rune(s)
		if len([]rune(best)) != len(runes) {
			t.Fatalf("rotation changed length for %q", s)
		}
		for r := 0; r < len(runes); r++ {
			rot := string(runes[r:]) + string(runes[:r])
			if rot < best {
				t.Fatalf("rotation %d of %q is smaller than %q", r, s, best)
			}
		}
		start := LeastRotation(s)
		got := string(runes[start:]) + string(runes[:start])
		if got != best {
			t.Fatalf("start index mismatch for %q: %q vs %q", s, got, best)
		}
		if LeastRotation(best) != 0 {
			t.Fatalf("canonical rotation of %q does not start at zero", best)
		}
	}
}
