package mod_fenwick_range_point_go

type FenwickTree struct {
	n    int
	tree []int64
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{n: n, tree: make([]int64, n+1)}
}

func (f *FenwickTree) Add(idx int, delta int64) {
	for i := idx + 1; i <= f.n; i += i & -i {
		f.tree[i] += delta
	}
}

func (f *FenwickTree) PrefixSum(idx int) int64 {
	var s int64
	for i := idx + 1; i > 0; i -= i & -i {
		s += f.tree[i]
	}
	return s
}

func (f *FenwickTree) KthSmallest(k int64) int {
	pos := 0
	pw := 1
	for pw<<1 <= f.n {
		pw <<= 1
	}
	for ; pw > 0; pw >>= 1 {
		if pos+pw <= f.n && f.tree[pos+pw] < k {
			pos += pw
			k -= f.tree[pos]
		}
	}
	return pos
}
