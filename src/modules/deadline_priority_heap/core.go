package mod_deadline_priority_heap

import "sync"

type DeadlineHeap struct {
	mu   sync.Mutex
	heap []*TimerTask
	byID map[string]*TimerTask
}

func NewDeadlineHeap() *DeadlineHeap {
	return &DeadlineHeap{byID: make(map[string]*TimerTask)}
}

func (d *DeadlineHeap) swap(i, j int) {
	d.heap[i], d.heap[j] = d.heap[j], d.heap[i]
	d.heap[i].index = i
	d.heap[j].index = j
}

func (d *DeadlineHeap) siftUp(i int) {
	for i > 0 {
		p := (i - 1) >> 1
		if d.heap[p].Deadline <= d.heap[i].Deadline {
			break
		}
		d.swap(p, i)
		i = p
	}
}

func (d *DeadlineHeap) siftDown(i int) {
	n := len(d.heap)
	for {
		l := 2*i + 1
		if l >= n {
			break
		}
		best := l
		if r := l + 1; r < n && d.heap[r].Deadline < d.heap[l].Deadline {
			best = r
		}
		if d.heap[i].Deadline <= d.heap[best].Deadline {
			break
		}
		d.swap(i, best)
		i = best
	}
}

func (d *DeadlineHeap) Schedule(id string, deadline int64) error {
	if id == "" || deadline < 0 {
		return ErrTimerHeapInvalid
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	if existing, ok := d.byID[id]; ok {
		old := existing.Deadline
		existing.Deadline = deadline
		if deadline < old {
			d.siftUp(existing.index)
		} else {
			d.siftDown(existing.index)
		}
		return nil
	}
	item := &TimerTask{ID: id, Deadline: deadline, index: len(d.heap)}
	d.heap = append(d.heap, item)
	d.byID[id] = item
	d.siftUp(item.index)
	return nil
}

func (d *DeadlineHeap) Cancel(id string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	item, ok := d.byID[id]
	if !ok {
		return false
	}
	idx := item.index
	last := len(d.heap) - 1
	delete(d.byID, id)
	if idx != last {
		d.swap(idx, last)
		d.heap = d.heap[:last]
		d.siftUp(idx)
		d.siftDown(idx)
	} else {
		d.heap = d.heap[:last]
	}
	return true
}

func (d *DeadlineHeap) DrainExpired(now int64) []string {
	d.mu.Lock()
	defer d.mu.Unlock()
	var expired []string
	for len(d.heap) > 0 && d.heap[0].Deadline <= now {
		root := d.heap[0]
		expired = append(expired, root.ID)
		delete(d.byID, root.ID)
		last := len(d.heap) - 1
		if last > 0 {
			d.heap[0] = d.heap[last]
			d.heap[0].index = 0
			d.heap = d.heap[:last]
			d.siftDown(0)
		} else {
			d.heap = d.heap[:0]
		}
	}
	return expired
}

func (d *DeadlineHeap) SnapshotDeadlines() []int64 {
	d.mu.Lock()
	defer d.mu.Unlock()
	out := make([]int64, len(d.heap))
	for i, item := range d.heap {
		out[i] = item.Deadline
	}
	return out
}
