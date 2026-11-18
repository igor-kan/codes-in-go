package mod_tarjan_scc_condense_go

import "testing"

func TestUnit_TarjanGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	p := TarjanSCC(4, [][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}})
	if p.Count != 2 || p.Comp[0] != p.Comp[2] {
		t.Fatalf("unexpected scc: %+v", p)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
