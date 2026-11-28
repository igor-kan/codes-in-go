package mod_boruvka_mst_go_d61

type DisjointSet struct {
	parent []int
	size   []int
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DisjointSet{parent: parent, size: size}
}

func (d *DisjointSet) Find(x int) int {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
}

func (d *DisjointSet) Union(a, b int) bool {
	ra := d.Find(a)
	rb := d.Find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}

func BoruvkaMST(n int, edges []WeightedEdge) MstSummary {
	if n <= 0 {
		return MstSummary{}
	}
	dsu := NewDisjointSet(n)
	var total int64
	used := 0
	for {
		best := make([]int, n)
		for i := range best {
			best[i] = -1
		}
		for idx, e := range edges {
			if e.U < 0 || e.U >= n || e.V < 0 || e.V >= n {
				continue
			}
			ru := dsu.Find(e.U)
			rv := dsu.Find(e.V)
			if ru == rv {
				continue
			}
			if best[ru] == -1 || e.W < edges[best[ru]].W {
				best[ru] = idx
			}
			if best[rv] == -1 || e.W < edges[best[rv]].W {
				best[rv] = idx
			}
		}
		merged := false
		for v := 0; v < n; v++ {
			idx := best[v]
			if idx == -1 {
				continue
			}
			e := edges[idx]
			if dsu.Union(e.U, e.V) {
				total += e.W
				used++
				merged = true
			}
		}
		if !merged {
			break
		}
	}
	return MstSummary{Weight: total, Edges: used}
}

type BoruvkaSolver struct{}

func (b BoruvkaSolver) MinimumSpanningTree(n int, edges []WeightedEdge) MstSummary {
	return BoruvkaMST(n, edges)
}

var _ MstSolver = BoruvkaSolver{}
