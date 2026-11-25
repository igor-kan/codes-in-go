package mod_consistent_ring_go_d60

import "testing"

func TestProp_HashRing(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	ring := NewConsistentRing(16)
	ring.AddNode("s1")
	ring.AddNode("s2")
	a, _ := ring.RouteKey("alpha")
	_a, _ := ring.RouteKey("alpha")
	if a != _a {
		t.Fatal("non-deterministic routing")
	}
}
