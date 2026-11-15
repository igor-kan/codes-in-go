package mod_vector_clock_causality

import "testing"

func TestVectorClockCausality(t *testing.T) {
	p0, _ := NewVectorClock(3)
	p1, _ := NewVectorClock(3)
	_ = p0.Tick(0)
	_ = p1.Tick(1)
	if rel, _ := p0.Compare(p1); rel != RelConcurrent {
		t.Fatalf("expected concurrent got %v", rel)
	}
	snap := p0.Clone()
	_ = p1.Merge(p0, 1)
	if !HappensBefore(snap, p1) {
		t.Fatal("expected causal send->recv happens-before")
	}
}
