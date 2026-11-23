package mod_circuit_halfopen_d58

import "testing"

func TestUnit_Breaker(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	cb := NewBreaker(2, 50)
	cb.OnFail(10); cb.OnFail(20)
	if cb.Allow(40) || !cb.Allow(75) {
		t.Fatalf("bad open/half-open timing")
	}
	cb.OnOk()
	if cb.state != StateClosed {
		t.Fatalf("expected closed after success")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
