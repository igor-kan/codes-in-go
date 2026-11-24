package mod_dijkstra_sp_go_d59

import "container/heap"

type dItem struct {
	d int64
	u int
}
type dHeap []dItem

func (h dHeap) Len() int           { return len(h) }
func (h dHeap) Less(i, j int) bool { return h[i].d < h[j].d }
func (h dHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *dHeap) Push(x any)        { *h = append(*h, x.(dItem)) }
func (h *dHeap) Pop() any {
	old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x
}

func DijkstraDist(n, src int, adj [][]WeightedArc) []int64 {
	const inf int64 = 1 << 60
	dist := make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0
	hp := &dHeap{{0, src}}
	heap.Init(hp)
	for hp.Len() > 0 {
		cur := heap.Pop(hp).(dItem)
		if cur.d > dist[cur.u] {
			continue
		}
		for _, e := range adj[cur.u] {
			nd := cur.d + e.Weight
			if nd < dist[e.To] {
				dist[e.To] = nd
				heap.Push(hp, dItem{nd, e.To})
			}
		}
	}
	return dist
}
