package mod_timing_wheel_scheduler

import "testing"

func TestTimingWheelMultiRound(t *testing.T) {
	tw, _ := NewTimingWheel(8)
	_, _ = tw.Schedule(1, 3, "t3")
	_, _ = tw.Schedule(2, 11, "t11") // slot 3, round 1!
	_, _ = tw.Schedule(3, 5, "cancel")
	if !tw.Cancel(3) {
		t.Fatal("expected cancel true")
	}
	f1 := tw.Advance(3)
	if len(f1) != 1 || f1[0].ID != 1 {
		t.Fatalf("tick 3 mismatch: %+v", f1)
	}
	f2 := tw.Advance(8) // tick 11
	if len(f2) != 1 || f2[0].ID != 2 || f2[0].Deadline != 11 {
		t.Fatalf("tick 11 mismatch: %+v", f2)
	}
	if err := VerifyTimingWheelBuckets(tw); err != nil {
		t.Fatal(err)
	}
}
