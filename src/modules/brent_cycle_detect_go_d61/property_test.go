package mod_brent_cycle_detect_go_d61

import "testing"

func TestProp_BrentCycle(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	for _, tc := range [][3]int{{3, 5, 17}, {2, 1, 31}, {7, 3, 64}} {
		a, bb, m := tc[0], tc[1], tc[2]
		f := func(x int) int { return (a*x + bb) % m }
		info := BrentCycle(f, 10%m)
		if !info.HasCycle || info.Lambda <= 0 || info.Mu < 0 {
			t.Fatalf("invalid cycle metadata: %+v", info)
		}
		x := 10 % m
		for i := 0; i < info.Mu; i++ {
			x = f(x)
		}
		y := x
		for i := 0; i < info.Lambda; i++ {
			y = f(y)
		}
		if x != y {
			t.Fatalf("mu/lambda do not reproduce the cycle: %+v", info)
		}
	}
}
