package mod_union_find_size_d57

import "testing"

func TestUnit_DsuSize(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	uf := NewUnionFindSize(5)
	uf.Union(0, 1)
	uf.Union(1, 2)
	if uf.Components() != 3 || uf.MaxSize() != 3 {
		t.Fatalf("bad union state")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
