package mod_hyperloglog_sketch_d57

import "testing"

func TestUnit_Hll(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	sk := NewHll64()
	if sk.Estimate() != 0.0 {
		t.Fatalf("empty estimate must be zero")
	}
	for x := uint64(0); x < 200; x++ {
		sk.Add(x)
	}
	est := sk.Estimate()
	if est < 100.0 || est > 350.0 {
		t.Fatalf("wildly off estimate: %v", est)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
