package mod_b_tree_order_m

import "testing"

func TestBTreeInsertSplitAndOrder(t *testing.T) {
	bt, err := NewBTree(3)
	if err != nil {
		t.Fatal(err)
	}
	for i := int64(50); i >= 1; i-- {
		bt.Put(i*7%100, i)
	}
	if err := VerifyBTreeInvariants(bt); err != nil {
		t.Fatal(err)
	}
	items := bt.InOrder()
	if len(items) != bt.Len() {
		t.Fatalf("len mismatch: %d != %d", len(items), bt.Len())
	}
}
