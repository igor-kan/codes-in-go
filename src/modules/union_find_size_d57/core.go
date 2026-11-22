package mod_union_find_size_d57

type UnionFindSize struct {
	parent []int
	sz     []int
	comps  int
	maxSz  int
}

func NewUnionFindSize(n int) *UnionFindSize {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	maxS := 0
	if n > 0 {
		maxS = 1
	}
	return &UnionFindSize{parent: p, sz: s, comps: n, maxSz: maxS}
}

func (u *UnionFindSize) Find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UnionFindSize) Union(a, b int) bool {
	ra, rb := u.Find(a), u.Find(b)
	if ra == rb {
		return false
	}
	if u.sz[ra] < u.sz[rb] {
		ra, rb = rb, ra
	}
	u.parent[rb] = ra
	u.sz[ra] += u.sz[rb]
	if u.sz[ra] > u.maxSz {
		u.maxSz = u.sz[ra]
	}
	u.comps--
	return true
}

func (u *UnionFindSize) MaxSize() int { return u.maxSz }
func (u *UnionFindSize) Components() int { return u.comps }
