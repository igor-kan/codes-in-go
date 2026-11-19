package mod_interval_tree_overlap_go

import "sort"

type ivNode struct {
	iv     ClosedIv
	maxEnd int64
	left   *ivNode
	right  *ivNode
}

type AugmentedIntervalTree struct {
	root *ivNode
}

func BuildIntervalTree(ivs []ClosedIv) *AugmentedIntervalTree {
	sorted := make([]ClosedIv, len(ivs))
	copy(sorted, ivs)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Start < sorted[j].Start })
	var build func(int, int) *ivNode
	build = func(l, r int) *ivNode {
		if l > r {
			return nil
		}
		mid := (l + r) >> 1
		ln := build(l, mid-1)
		rn := build(mid+1, r)
		mx := sorted[mid].End
		if ln != nil && ln.maxEnd > mx {
			mx = ln.maxEnd
		}
		if rn != nil && rn.maxEnd > mx {
			mx = rn.maxEnd
		}
		return &ivNode{iv: sorted[mid], maxEnd: mx, left: ln, right: rn}
	}
	return &AugmentedIntervalTree{root: build(0, len(sorted)-1)}
}

func (t *AugmentedIntervalTree) QueryOverlap(lo, hi int64) []ClosedIv {
	var out []ClosedIv
	var visit func(*ivNode)
	visit = func(nd *ivNode) {
		if nd == nil || nd.maxEnd < lo {
			return
		}
		visit(nd.left)
		if nd.iv.Start <= hi && nd.iv.End >= lo {
			out = append(out, nd.iv)
		}
		if nd.iv.Start <= hi {
			visit(nd.right)
		}
	}
	visit(t.root)
	return out
}
