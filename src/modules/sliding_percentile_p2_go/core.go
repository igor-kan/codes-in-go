package mod_sliding_percentile_p2_go

import "sort"

type SlidingQuantileWindow struct {
	buf  []float64
	cap  int
	head int
}

func NewSlidingQuantileWindow(capacity int) *SlidingQuantileWindow {
	return &SlidingQuantileWindow{cap: capacity}
}

func (w *SlidingQuantileWindow) Push(x float64) {
	if len(w.buf) < w.cap {
		w.buf = append(w.buf, x)
	} else {
		w.buf[w.head] = x
		w.head = (w.head + 1) % w.cap
	}
}

func (w *SlidingQuantileWindow) Summary() WindowSummary {
	cp := make([]float64, len(w.buf))
	copy(cp, w.buf)
	sort.Float64s(cp)
	n := len(cp)
	med := cp[n/2]
	if n%2 == 0 {
		med = 0.5 * (cp[n/2-1] + cp[n/2])
	}
	return WindowSummary{Min: cp[0], Median: med, Max: cp[n-1]}
}
