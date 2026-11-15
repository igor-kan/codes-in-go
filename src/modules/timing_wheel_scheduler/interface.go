package mod_timing_wheel_scheduler

import "container/list"

// NewTimingWheel constructs a wheel with numSlots (>0) circular buckets.
func NewTimingWheel(numSlots int) (*TimingWheel, error) {
	if numSlots <= 0 {
		return nil, ErrTimerInvalid
	}
	buckets := make([]*list.List, numSlots)
	for i := range buckets {
		buckets[i] = list.New()
	}
	return &TimingWheel{
		NumSlots: numSlots,
		Buckets:  buckets,
		Active:   make(map[uint64]*TimerEntry),
	}, nil
}
