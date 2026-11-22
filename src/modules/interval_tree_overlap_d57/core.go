package mod_interval_tree_overlap_d57

import "sort"

func MaxIntervalOverlap(spans []TimeSpan) int {
	if len(spans) == 0 {
		return 0
	}
	type evt struct {
		t int64
		d int
	}
	evts := make([]evt, 0, len(spans)*2)
	for _, sp := range spans {
		if sp.End <= sp.Start {
			continue
		}
		evts = append(evts, evt{sp.Start, 1}, evt{sp.End, -1})
	}
	sort.Slice(evts, func(i, j int) bool {
		if evts[i].t != evts[j].t {
			return evts[i].t < evts[j].t
		}
		return evts[i].d < evts[j].d
	})
	cur, best := 0, 0
	for _, e := range evts {
		cur += e.d
		if cur > best {
			best = cur
		}
	}
	return best
}
