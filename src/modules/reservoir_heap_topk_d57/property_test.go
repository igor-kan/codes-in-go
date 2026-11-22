package mod_reservoir_heap_topk_d57

import "testing"

func TestProp_TopK(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	tk := NewTopK(2)
	tk.Push(-10); tk.Push(-3); tk.Push(-20)
	res := tk.TopDesc()
	if res[0] != -3 || res[1] != -10 {
		t.Fatalf("negative topk failed: %v", res)
	}
}
