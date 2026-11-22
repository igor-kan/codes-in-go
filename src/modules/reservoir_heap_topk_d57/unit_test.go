package mod_reservoir_heap_topk_d57

import "testing"

func TestUnit_TopK(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	tk := NewTopK(3)
	for _, v := range []int64{4, 19, 2, 11, 25, 8} {
		tk.Push(v)
	}
	res := tk.TopDesc()
	if len(res) != 3 || res[0] != 25 || res[1] != 19 || res[2] != 11 {
		t.Fatalf("bad topk: %v", res)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
