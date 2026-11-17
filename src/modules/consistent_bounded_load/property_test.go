package mod_consistent_bounded_load

import "testing"

func TestBoundedLoadConcurrent(t *testing.T) {
	ring, _ := NewBoundedHashRing([]string{"s1", "s2", "s3", "s4"}, 12, 1.5)
	keys := make([]string, 80)
	for i := range keys {
		keys[i] = "req"
	}
	srvs := parallelAcquireBatch(ring, keys)
	if err := verifyBoundedLoad(ring); err != nil {
		t.Fatal(err)
	}
	for _, s := range srvs {
		if err := ring.Release(s); err != nil {
			t.Fatal(err)
		}
	}
	_, tot := ring.LoadSnapshot()
	if tot != 0 {
		t.Fatalf("expected 0 active requests after release, got %d", tot)
	}
}
