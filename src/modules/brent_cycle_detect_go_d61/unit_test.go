package mod_brent_cycle_detect_go_d61

import "testing"

func TestUnit_BrentCycle(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	next := []int{1, 2, 3, 4, 5, 6, 7, 5}
	f := func(x int) int { return next[x] }
	info := BrentCycle(f, 0)
	if info.Mu != 5 || info.Lambda != 3 || !info.HasCycle {
		t.Fatalf("unexpected cycle info: %+v", info)
	}
	step := func(x int) int { return (x + 1) % 10 }
	loop := BrentCycle(step, 0)
	if loop.Mu != 0 || loop.Lambda != 10 {
		t.Fatalf("pure cycle misdetected: %+v", loop)
	}
	d := BrentDetector{}
	if d.Detect(step, 0).Lambda != 10 {
		t.Fatal("detector wrapper mismatch")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
