package mod_union_find_rollback_d60

type DsuWithRollback struct {
	parent []int
	rank   []int
	hist   []UndoEntry
	comps  int
}

func NewDsuWithRollback(n int) *DsuWithRollback {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DsuWithRollback{parent: p, rank: make([]int, n), comps: n}
}

func (d *DsuWithRollback) Find(x int) int {
	for d.parent[x] != x {
		x = d.parent[x]
	}
	return x
}

func (d *DsuWithRollback) Snapshot() int { return len(d.hist) }

func (d *DsuWithRollback) Components() int { return d.comps }

func (d *DsuWithRollback) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}
	if d.rank[ra] < d.rank[rb] {
		ra, rb = rb, ra
	}
	d.hist = append(d.hist, UndoEntry{Child: rb, Root: ra, Rank: d.rank[ra]})
	d.parent[rb] = ra
	if d.rank[ra] == d.rank[rb] {
		d.rank[ra]++
	}
	d.comps--
	return true
}

func (d *DsuWithRollback) Rollback(snap int) {
	for len(d.hist) > snap {
		top := d.hist[len(d.hist)-1]
		d.hist = d.hist[:len(d.hist)-1]
		d.parent[top.Child] = top.Child
		d.rank[top.Root] = top.Rank
		d.comps++
	}
}
