package mod_segment_tree_concurrent

import "testing"

func TestSegTreeUpdateAndQuery(t *testing.T) {
	raw := []int64{3, 1, 4, 1, 5, 9, 2, 6}
	st, err := NewConcurrentSegTree(raw)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifySegMatchesNaive(st, raw, 1, 6); err != nil {
		t.Fatal(err)
	}
	_ = st.Update(3, 20)
	raw[3] = 20
	if err := verifySegMatchesNaive(st, raw, 0, 8); err != nil {
		t.Fatal(err)
	}
}
