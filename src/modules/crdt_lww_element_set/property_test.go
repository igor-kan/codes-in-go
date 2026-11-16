package mod_crdt_lww_element_set

import "testing"

func TestLWWAlgebraicLawsAndConcurrency(t *testing.T) {
	a := NewLWWSet(BiasRemoveWins)
	b := NewLWWSet(BiasRemoveWins)
	c := NewLWWSet(BiasRemoveWins)
	_ = a.Add("k1", 5); _ = a.Remove("k2", 7)
	_ = b.Add("k2", 7); _ = b.Add("k3", 12)
	_ = c.Remove("k3", 11); _ = c.Add("k4", 3)
	if err := VerifyLWWLaws(a, b, c); err != nil {
		t.Fatalf("CRDT laws failed: %v", err)
	}
	m, err := ConvergeReplicas([]*LWWSet{a, b, c})
	if err != nil || len(m.Elements()) != 3 {
		t.Fatalf("unexpected converged elements: %v err=%v", m.Elements(), err)
	}
}
