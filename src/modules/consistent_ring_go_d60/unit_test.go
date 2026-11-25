package mod_consistent_ring_go_d60

import "testing"

func TestUnit_HashRing(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	ring := NewConsistentRing(12)
	ring.AddNode("node-a")
	ring.AddNode("node-b")
	ring.AddNode("node-c")
	n1, ok := ring.RouteKey("user:42")
	if !ok || n1 == "" {
		t.Fatal("expected valid shard")
	}
	ring.RemoveNode("node-b")
	n2, ok2 := ring.RouteKey("user:42")
	if !ok2 || n2 == "node-b" {
		t.Fatal("removed node still routed")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
