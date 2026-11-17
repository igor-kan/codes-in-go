package mod_deadline_priority_heap

func verifyHeapOrder(h *DeadlineHeap) error {
	dl := h.SnapshotDeadlines()
	for i := 1; i < len(dl); i++ {
		p := (i - 1) >> 1
		if dl[p] > dl[i] {
			return ErrTimerHeapInvariant
		}
	}
	return nil
}
