package mod_token_bucket_limiter_d56

import "testing"

func TestUnit_Limiter(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	tb := NewTokenBucket(5.0, 2.0)
	if !tb.Allow(0.0, 4.0) || tb.Allow(0.0, 2.0) {
		t.Fatalf("unexpected initial admission")
	}
	if !tb.Allow(1.0, 2.0) {
		t.Fatalf("expected refill by t=1.0")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
