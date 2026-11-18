package mod_weighted_round_robin_gcd

import "testing"

func TestUnit_SmoothWrr(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	b := NewSmoothWRR([]int{5, 1, 1})
	counts := make([]int, 3)
	first := b.NextPeer()
	second := b.NextPeer()
	counts[first]++
	counts[second]++
	for i := 2; i < 7; i++ {
		counts[b.NextPeer()]++
	}
	if counts[0] != 5 || counts[1] != 1 || counts[2] != 1 || first != 0 {
		t.Fatalf("unexpected distribution: %v first=%d", counts, first)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
