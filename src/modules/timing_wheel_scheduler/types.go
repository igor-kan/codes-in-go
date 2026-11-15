package mod_timing_wheel_scheduler

import "container/list"

// TimerEntry represents a scheduled timer inside a hashed wheel slot.
type TimerEntry struct {
	ID       uint64
	Rounds   uint64
	Deadline uint64
	Payload  string
	slotIdx  int
	elem     *list.Element
	active   bool
}

// TimingWheel manages fixed-slot circular timer buckets.
type TimingWheel struct {
	NumSlots    int
	CurrentTick uint64
	Buckets     []*list.List
	Active      map[uint64]*TimerEntry
}
