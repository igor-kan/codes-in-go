package mod_singleflight_dedup_group

import "testing"

func TestSingleflightCoalescesBurst(t *testing.T) {
	execs, vals := RunStampedBurst("hot_key_1", 16)
	if err := VerifySingleExecution(execs, vals); err != nil {
		t.Fatalf("expected 1 execution across 16 concurrent callers, got execs=%d err=%v", execs, err)
	}
}
