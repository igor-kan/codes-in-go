package mod_union_find_rollback_d60

import "testing"

func TestUnit_RollbackDsu(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	d := NewDsuWithRollback(5)
	d.Union(0, 1)
	snap := d.Snapshot()
	d.Union(1, 2)
	if d.Find(0) != d.Find(2) || d.Components() != 3 {
		t.Fatal("union failed")
	}
	d.Rollback(snap)
	if d.Find(0) == d.Find(2) || d.Components() != 4 {
		t.Fatal("rollback failed")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
