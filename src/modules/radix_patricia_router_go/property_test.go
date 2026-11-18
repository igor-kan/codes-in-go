package mod_radix_patricia_router_go

import "testing"

func TestProp_RadixRouter(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	rt := NewBytePrefixRouter()
	rt.Insert("/a", 10)
	if rt.MatchLongest("/b").Found {
		t.Fatal("unexpected match")
	}
}
