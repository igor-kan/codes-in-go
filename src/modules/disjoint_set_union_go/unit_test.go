package mod_disjoint_set_union_go

import "testing"

func TestUnit_DsuGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	d := NewDisjointSet(5)
	d.Union(0, 1); d.Union(2, 3); d.Union(1, 2)
	if d.Components() != 2 || d.Find(0) != d.Find(3) || d.Union(0, 3) {
		t.Fatal("DSU merge mismatch")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
