package mod_deadline_priority_heap

import "errors"

var (
	ErrTimerHeapInvalid   = errors.New("timerheap: invalid argument")
	ErrTimerHeapBounds    = errors.New("timerheap: index out of bounds")
	ErrTimerHeapEmpty     = errors.New("timerheap: empty container")
	ErrTimerHeapInvariant = errors.New("timerheap: invariant violated")
)
