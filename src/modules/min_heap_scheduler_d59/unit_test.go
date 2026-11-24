package mod_min_heap_scheduler_d59

import "testing"

func TestUnit_Sched(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	ord := RunByDeadline([]TaskItem{{30, 1, 3}, {10, 2, 1}, {10, 1, 0}, {20, 1, 2}})
	if len(ord) != 4 || ord[0] != 0 || ord[1] != 1 || ord[2] != 2 || ord[3] != 3 {
		t.Fatalf("bad schedule order: %v", ord)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
