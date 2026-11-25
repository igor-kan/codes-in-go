package mod_hyperloglog_card_d60

import "testing"

func TestUnit_HllEstimator(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	hll := NewHyperLogLog(6)
	for i := uint64(1); i <= 200; i++ {
		hll.Observe(i)
		hll.Observe(i)
	}
	c := hll.Count()
	if c < 130 || c > 280 {
		t.Fatalf("unexpected HLL estimate %d for 200 distinct items", c)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
