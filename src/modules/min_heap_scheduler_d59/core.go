package mod_min_heap_scheduler_d59

import "container/heap"

type taskHeap []TaskItem

func (h taskHeap) Len() int { return len(h) }
func (h taskHeap) Less(i, j int) bool {
	if h[i].Deadline != h[j].Deadline {
		return h[i].Deadline < h[j].Deadline
	}
	return h[i].Seq < h[j].Seq
}
func (h taskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *taskHeap) Push(x any)   { *h = append(*h, x.(TaskItem)) }
func (h *taskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func RunByDeadline(items []TaskItem) []int {
	hp := append(taskHeap(nil), items...)
	heap.Init(&hp)
	out := make([]int, 0, len(items))
	for hp.Len() > 0 {
		out = append(out, heap.Pop(&hp).(TaskItem).Payload)
	}
	return out
}
