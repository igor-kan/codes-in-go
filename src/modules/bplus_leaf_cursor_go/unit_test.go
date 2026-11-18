package mod_bplus_leaf_cursor_go

import "testing"

func TestUnit_BPlusCursor(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	idx := NewBPlusLeafIndex()
	for _, k := range []int64{40, 10, 30, 20, 50} {
		idx.Put(k, k*10)
	}
	res := idx.Scan(15, 42)
	if len(res) != 3 || res[0].Key != 20 || res[2].Key != 40 {
		t.Fatalf("unexpected scan: %v", res)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
