package mod_timing_wheel_scheduler

import "testing"

func TestTimingWheelDeadlinesExact(t *testing.T) {
	tw, _ := NewTimingWheel(16)
	for id := uint64(1); id <= 40; id++ {
		delay := (id*7)%35 + 1
		_, _ = tw.Schedule(id, delay, "job")
	}
	for tick := uint64(1); tick <= 40; tick++ {
		fired := tw.Advance(1)
		for _, f := range fired {
			if f.Deadline != tick {
				t.Fatalf("fired at tick %d with deadline %d", tick, f.Deadline)
			}
		}
	}
	if tw.PendingCount() != 0 || VerifyTimingWheelBuckets(tw) != nil {
		t.Fatal("timers leaked")
	}
}
