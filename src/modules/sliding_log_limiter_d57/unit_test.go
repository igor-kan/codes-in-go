package mod_sliding_log_limiter_d57

import "testing"

func TestUnit_SLog(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	sl := NewSlidingLog(2, 10)
	if !sl.Allow(1) || !sl.Allow(5) || sl.Allow(9) {
		t.Fatalf("bad burst admission")
	}
	if !sl.Allow(12) {
		t.Fatalf("expected old event at t=1 expired by t=12")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
