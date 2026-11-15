package mod_swiss_probe_hashmap

import "testing"

func TestSwissDifferentialAgainstGoMap(t *testing.T) {
	sw, _ := NewSwissMap(1)
	ref := make(map[uint64]uint64)
	for step := uint64(0); step < 400; step++ {
		k := (step*37 + 11) % 73
		if step%4 == 0 {
			sw.Delete(k)
			delete(ref, k)
		} else {
			sw.Put(k, step)
			ref[k] = step
		}
		if sw.Count != len(ref) {
			t.Fatalf("count mismatch: got %d want %d", sw.Count, len(ref))
		}
	}
	for k, want := range ref {
		if got, ok := sw.Get(k); !ok || got != want {
			t.Fatalf("key %d got %v (%v) want %v", k, got, ok, want)
		}
	}
}
