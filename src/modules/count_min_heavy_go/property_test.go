package mod_count_min_heavy_go

import "testing"

func TestCmsConcurrentIngest(t *testing.T) {
	cms, _ := NewConservativeCMS(4, 64)
	b1 := []string{"x", "x", "x", "y"}
	b2 := []string{"x", "y", "z", "x"}
	parallelIngest(cms, [][]string{b1, b2})
	if err := verifyNoUnderestimate(cms, map[string]uint64{"x": 5, "y": 2, "z": 1}); err != nil {
		t.Fatal(err)
	}
}
