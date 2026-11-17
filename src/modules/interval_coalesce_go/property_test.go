package mod_interval_coalesce_go

import "testing"

func TestDisjointParallelProbes(t *testing.T) {
	l := NewDisjointIntervals()
	_ = l.Add(0, 10)
	_ = l.Add(20, 30)
	hits := parallelProbePoints(l, []int64{5, 15, 25, 31})
	if !hits[0] || hits[1] || !hits[2] || hits[3] {
		t.Fatalf("unexpected probe hits: %v", hits)
	}
}
