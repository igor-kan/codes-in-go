package mod_deadline_priority_heap

import "testing"

func TestDeadlineCancelRescheduleDrain(t *testing.T) {
	h := NewDeadlineHeap()
	_ = h.Schedule("a", 50)
	_ = h.Schedule("b", 20)
	_ = h.Schedule("c", 30)
	_ = h.Schedule("a", 10) // move "a" earliest
	if !h.Cancel("b") {
		t.Fatal("expected cancel of b")
	}
	if err := verifyHeapOrder(h); err != nil {
		t.Fatal(err)
	}
	if got := fastCountDueBefore(h.SnapshotDeadlines(), 25); got != 1 {
		t.Fatalf("expected 1 due item, got %d", got)
	}
	exp := h.DrainExpired(35)
	if len(exp) != 2 || exp[0] != "a" || exp[1] != "c" {
		t.Fatalf("unexpected drain order: %v", exp)
	}
}
