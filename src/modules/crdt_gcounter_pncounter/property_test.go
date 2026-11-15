package mod_crdt_gcounter_pncounter

import "testing"

func TestCRDTLatticeAxioms(t *testing.T) {
	a, _ := NewPNCounter(4)
	b, _ := NewPNCounter(4)
	c, _ := NewPNCounter(4)
	_ = a.Add(0, 5)
	_ = b.Add(1, -2)
	_ = b.Add(0, 3)
	_ = c.Add(3, 9)
	if err := VerifySemilatticeLaws(a, b, c); err != nil {
		t.Fatal(err)
	}
}
