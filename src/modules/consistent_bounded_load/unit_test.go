package mod_consistent_bounded_load

import "testing"

func TestBoundedLoadSpillover(t *testing.T) {
	ring, err := NewBoundedHashRing([]string{"node-a", "node-b", "node-c"}, 8, 1.25)
	if err != nil {
		t.Fatal(err)
	}
	// Acquire 60 requests with identical key: must spill across all servers!
	for i := 0; i < 60; i++ {
		if _, err := ring.Acquire("hot-key"); err != nil {
			t.Fatal(err)
		}
	}
	if err := verifyBoundedLoad(ring); err != nil {
		t.Fatal(err)
	}
	loads, _ := ring.LoadSnapshot()
	if loads["node-a"] == 60 || loads["node-b"] == 60 || loads["node-c"] == 60 {
		t.Fatalf("hot key was not balanced: %v", loads)
	}
}
