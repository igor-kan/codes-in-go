package mod_deadline_priority_heap

func validateDeadline(deadline int64) error {
	if deadline < 0 {
		return ErrTimerHeapInvalid
	}
	return nil
}
