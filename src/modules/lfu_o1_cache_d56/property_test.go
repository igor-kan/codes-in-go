package mod_lfu_o1_cache_d56

import "testing"

func TestProp_Lfu(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	c := NewLfuCache(3)
	for i := 0; i < 20; i++ {
		c.Put(i%5, i)
		if len(c.vals) > 3 {
			t.Fatalf("capacity exceeded")
		}
	}
}
