package mod_disjoint_set_union_go

type DisjointSet struct {
	parent     []int
	size       []int
	components int
}

func NewDisjointSet(n int) *DisjointSet {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &DisjointSet{parent: p, size: s, components: n}
}

func (d *DisjointSet) Find(x int) int {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
}

func (d *DisjointSet) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	d.components--
	return true
}

func (d *DisjointSet) Components() int {
	return d.components
}
