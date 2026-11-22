package mod_reservoir_heap_topk_d57

import (
	"container/heap"
	"sort"
)

type int64MinHeap []int64

func (h int64MinHeap) Len() int           { return len(h) }
func (h int64MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h int64MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *int64MinHeap) Push(x any)        { *h = append(*h, x.(int64)) }
func (h *int64MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type TopKStream struct {
	k int
	h int64MinHeap
}

func NewTopK(k int) *TopKStream {
	return &TopKStream{k: k}
}

func (tk *TopKStream) Push(v int64) {
	if tk.k <= 0 {
		return
	}
	if len(tk.h) < tk.k {
		heap.Push(&tk.h, v)
	} else if v > tk.h[0] {
		tk.h[0] = v
		heap.Fix(&tk.h, 0)
	}
}

func (tk *TopKStream) TopDesc() []int64 {
	out := append([]int64(nil), tk.h...)
	sort.Slice(out, func(i, j int) bool { return out[i] > out[j] })
	return out
}
