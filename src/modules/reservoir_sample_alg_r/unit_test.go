package mod_reservoir_sample_alg_r

import "testing"

func TestUnit_ReservoirGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	s1 := NewReservoirSampler(5, 42)
	s2 := NewReservoirSampler(5, 42)
	for x := int64(1); x <= 100; x++ {
		s1.Offer(x); s2.Offer(x)
	}
	i1, i2 := s1.Items(), s2.Items()
	if len(i1) != 5 || i1[0] != i2[0] || i1[4] != i2[4] {
		t.Fatal("determinism failed")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
