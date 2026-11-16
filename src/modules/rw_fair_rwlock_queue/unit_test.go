package mod_rw_fair_rwlock_queue

import "testing"

func TestFairRWMutualExclusion(t *testing.T) {
	v, err := StressFairRW(6, 3, 200)
	if err != nil {
		t.Fatalf("fair rw stress failed: %v", err)
	}
	if err := VerifyMutualExclusion(3, 200, v); err != nil {
		t.Fatal(err)
	}
}
