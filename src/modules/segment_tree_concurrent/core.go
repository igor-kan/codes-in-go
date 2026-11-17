package mod_segment_tree_concurrent

import (
	"math"
	"sync"
)

type ConcurrentSegTree struct {
	mu   sync.RWMutex
	n    int
	tree []SegStats
}

func mergeStats(a, b SegStats) SegStats {
	mx := a.Max
	if b.Max > mx {
		mx = b.Max
	}
	return SegStats{Sum: a.Sum + b.Sum, Max: mx}
}

func NewConcurrentSegTree(initVals []int64) (*ConcurrentSegTree, error) {
	if len(initVals) == 0 {
		return nil, ErrSegTreeEmpty
	}
	n := len(initVals)
	t := make([]SegStats, 2*n)
	for i := 0; i < n; i++ {
		t[n+i] = SegStats{Sum: initVals[i], Max: initVals[i]}
	}
	for i := n - 1; i > 0; i-- {
		t[i] = mergeStats(t[2*i], t[2*i+1])
	}
	return &ConcurrentSegTree{n: n, tree: t}, nil
}

func (s *ConcurrentSegTree) Update(idx int, val int64) error {
	if idx < 0 || idx >= s.n {
		return ErrSegTreeBounds
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	p := idx + s.n
	s.tree[p] = SegStats{Sum: val, Max: val}
	for p >>= 1; p > 0; p >>= 1 {
		s.tree[p] = mergeStats(s.tree[2*p], s.tree[2*p+1])
	}
	return nil
}

func (s *ConcurrentSegTree) Query(l, r int) (SegStats, error) {
	if err := validateSegRange(s.n, l, r); err != nil {
		return SegStats{}, err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	accL := SegStats{Sum: 0, Max: math.MinInt64}
	accR := SegStats{Sum: 0, Max: math.MinInt64}
	l += s.n
	r += s.n
	for ; l < r; l, r = l>>1, r>>1 {
		if l&1 == 1 {
			accL = mergeStats(accL, s.tree[l])
			l++
		}
		if r&1 == 1 {
			r--
			accR = mergeStats(s.tree[r], accR)
		}
	}
	return mergeStats(accL, accR), nil
}
