package mod_interval_coalesce_go

import (
	"sort"
	"sync"
)

type DisjointIntervals struct {
	mu    sync.RWMutex
	spans []Span
}

func NewDisjointIntervals() *DisjointIntervals {
	return &DisjointIntervals{}
}

func (d *DisjointIntervals) Add(start, end int64) error {
	if start > end {
		return ErrIntervalInvalid
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	cur := Span{Start: start, End: end}
	next := make([]Span, 0, len(d.spans)+1)
	inserted := false
	for _, sp := range d.spans {
		if sp.End+1 < cur.Start {
			next = append(next, sp)
		} else if cur.End+1 < sp.Start {
			if !inserted {
				next = append(next, cur)
				inserted = true
			}
			next = append(next, sp)
		} else {
			if sp.Start < cur.Start {
				cur.Start = sp.Start
			}
			if sp.End > cur.End {
				cur.End = sp.End
			}
		}
	}
	if !inserted {
		next = append(next, cur)
	}
	d.spans = next
	return nil
}

func (d *DisjointIntervals) Contains(x int64) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	idx := sort.Search(len(d.spans), func(i int) bool { return d.spans[i].End >= x })
	return idx < len(d.spans) && d.spans[idx].Start <= x
}

func (d *DisjointIntervals) Snapshots() []Span {
	d.mu.RLock()
	defer d.mu.RUnlock()
	cp := make([]Span, len(d.spans))
	copy(cp, d.spans)
	return cp
}
