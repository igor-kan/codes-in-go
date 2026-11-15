package mod_crdt_gcounter_pncounter

import "testing"

func TestCRDTConvergence(t *testing.T) {
	r0, _ := NewPNCounter(3)
	r1, _ := NewPNCounter(3)
	_ = r0.Add(0, 10)
	_ = r0.Add(0, -3)
	_ = r1.Add(1, 5)
	_ = r1.Add(1, -8)
	_ = r0.Merge(r1)
	if r0.Value() != 4 {
		t.Fatalf("got %d want 4", r0.Value())
	}
}
