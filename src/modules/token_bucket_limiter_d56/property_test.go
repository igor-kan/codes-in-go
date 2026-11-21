package mod_token_bucket_limiter_d56

import "testing"

func TestProp_Limiter(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	tb := NewTokenBucket(3.0, 1.0)
	_ = tb.Allow(100.0, 0.0)
	if tb.tokens > 3.000001 || tb.tokens < 0 {
		t.Fatalf("tokens out of bounds: %v", tb.tokens)
	}
}
