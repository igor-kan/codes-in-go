package mod_interval_coalesce_go

import "testing"

func TestDisjointCoalescing(t *testing.T) {
	l := NewDisjointIntervals()
	_ = l.Add(10, 15)
	_ = l.Add(1, 5)
	_ = l.Add(6, 9) // bridges [1,5] and [10,15] into [1,15]
	sp := l.Snapshots()
	if err := verifyStrictDisjoint(sp); err != nil {
		t.Fatal(err)
	}
	if len(sp) != 1 || sp[0].Start != 1 || sp[0].End != 15 {
		t.Fatalf("expected merged [1,15], got %v", sp)
	}
	if fastTotalCovered(sp) != 15 {
		t.Fatalf("expected 15 covered points")
	}
}
