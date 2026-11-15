package mod_vector_clock_causality

import "testing"

func TestVectorClockTransitiveChain(t *testing.T) {
	a, _ := NewVectorClock(3)
	_ = a.Tick(0)
	b := a.Clone()
	_ = b.Merge(a, 1)
	c := b.Clone()
	_ = c.Merge(b, 2)
	if err := VerifyTransitivity(a, b, c); err != nil {
		t.Fatal(err)
	}
}
