package mod_sliding_log_limiter_d57

type SlidingLog struct {
	limit  int
	window int64
	events []int64
}

func NewSlidingLog(limit int, window int64) *SlidingLog {
	return &SlidingLog{limit: limit, window: window}
}

func (s *SlidingLog) Allow(now int64) bool {
	cutoff := now - s.window
	drop := 0
	for drop < len(s.events) && s.events[drop] <= cutoff {
		drop++
	}
	if drop > 0 {
		s.events = append([]int64(nil), s.events[drop:]...)
	}
	if len(s.events) >= s.limit {
		return false
	}
	s.events = append(s.events, now)
	return true
}
