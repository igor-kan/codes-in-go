package mod_deadline_priority_heap

import "testing"

func TestDeadlineConcurrentSchedule(t *testing.T) {
	h := NewDeadlineHeap()
	parallelScheduleBatch(h, 40)
	if err := verifyHeapOrder(h); err != nil {
		t.Fatal(err)
	}
	exp := h.DrainExpired(1000)
	if len(exp) != 40 {
		t.Fatalf("expected all 40 drained, got %d", len(exp))
	}
}
