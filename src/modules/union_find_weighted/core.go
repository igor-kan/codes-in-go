package mod_union_find_weighted

import "math"

// Find returns root of x while path-compressing and updating Weight[x] = val[x] - val[root].
func (d *WeightedDSU) Find(x int) int {
	if d.Parent[x] == x {
		return x
	}
	r := d.Find(d.Parent[x])
	d.Weight[x] += d.Weight[d.Parent[x]]
	d.Parent[x] = r
	return r
}

// Relate asserts val[x] - val[y] = diff, returning (consistent, err).
func (d *WeightedDSU) Relate(x, y int, diff float64) (bool, error) {
	if x < 0 || y < 0 || x >= len(d.Parent) || y >= len(d.Parent) {
		return false, ErrWDSUBounds
	}
	rx := d.Find(x)
	ry := d.Find(y)
	if rx == ry {
		actual := d.Weight[x] - d.Weight[y]
		return math.Abs(actual-diff) <= 1e-9, nil
	}
	// Attach rx under ry: val[x]-val[y] = diff => (val[rx]+wx) - (val[ry]+wy) = diff => val[rx]-val[ry] = diff + wy - wx
	if d.Rank[rx] < d.Rank[ry] {
		d.Parent[rx] = ry
		d.Weight[rx] = diff + d.Weight[y] - d.Weight[x]
	} else {
		d.Parent[ry] = rx
		d.Weight[ry] = d.Weight[x] - d.Weight[y] - diff
		if d.Rank[rx] == d.Rank[ry] {
			d.Rank[rx]++
		}
	}
	return true, nil
}

// Diff returns (val[x] - val[y], true) iff x and y lie in the same component.
func (d *WeightedDSU) Diff(x, y int) (float64, bool) {
	if d.Find(x) != d.Find(y) {
		return 0, false
	}
	return d.Weight[x] - d.Weight[y], true
}
