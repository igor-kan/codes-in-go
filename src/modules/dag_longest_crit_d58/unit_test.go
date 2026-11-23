package mod_dag_longest_crit_d58

import "testing"

func TestUnit_Crit(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	ms, ok := DagMakespan(4, []TaskEdge{{0, 1, 3}, {0, 2, 7}, {1, 3, 2}, {2, 3, 4}})
	if !ok || ms != 11 {
		t.Fatalf("expected 11, got %d %v", ms, ok)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
