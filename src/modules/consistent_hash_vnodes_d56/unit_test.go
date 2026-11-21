package mod_consistent_hash_vnodes_d56

import "testing"

func TestUnit_Ring(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	hr := NewHashRing([]string{"node-a", "node-b", "node-c"}, 16)
	h1 := hr.Lookup("session:12345")
	h2 := hr.Lookup("session:12345")
	if h1 == "" || h1 != h2 {
		t.Fatalf("non-deterministic routing: %q vs %q", h1, h2)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
