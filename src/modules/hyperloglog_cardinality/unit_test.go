package mod_hyperloglog_cardinality

import (
	"math"
	"testing"
)

func TestHLLAccuracy(t *testing.T) {
	h, err := NewHLL(12, 0x1234)
	if err != nil {
		t.Fatal(err)
	}
	for i := uint64(0); i < 6000; i++ {
		h.Add(i)
		h.Add(i)
	}
	if err := VerifyHLLMonotone(h); err != nil {
		t.Fatal(err)
	}
	est := h.Count()
	if math.Abs(est-6000.0)/6000.0 > 0.08 {
		t.Fatalf("excessive relative error: got %v want ~6000", est)
	}
}
