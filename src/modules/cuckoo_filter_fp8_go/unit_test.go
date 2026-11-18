package mod_cuckoo_filter_fp8_go

import "testing"

func TestUnit_CuckooGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	cf := NewCuckooFilter(64)
	for k := uint64(1); k <= 30; k++ {
		if !cf.Insert(k) || !cf.Lookup(k) {
			t.Fatalf("failed on key %d", k)
		}
	}
	cf.Delete(15)
	if cf.Lookup(15) {
		t.Fatal("expected key 15 deleted")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
