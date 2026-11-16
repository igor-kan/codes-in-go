package mod_reservoir_quantiles_gk

import "math"

func NewGKSummary(eps float64) (*GKSummary, error) {
	if err := ValidateEps(eps); err != nil {
		return nil, err
	}
	return &GKSummary{
		eps: eps,
	}, nil
}

func (g *GKSummary) Observe(v float64) {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return
	}
	threshold := int(math.Floor(2.0 * g.eps * float64(g.n)))
	if g.n == 0 {
		g.tuples = append(g.tuples, gkTuple{v: v, g: 1, delta: 0})
		g.n = 1
		return
	}
	pos := 0
	for pos < len(g.tuples) && g.tuples[pos].v < v {
		pos++
	}
	delta := 0
	if pos > 0 && pos < len(g.tuples) {
		delta = threshold
		if delta < 0 {
			delta = 0
		}
	}
	g.tuples = append(g.tuples, gkTuple{})
	copy(g.tuples[pos+1:], g.tuples[pos:])
	g.tuples[pos] = gkTuple{v: v, g: 1, delta: delta}
	g.n++
	g.band++
	if g.band > int(math.Ceil(1.0/(2.0*g.eps)))+1 {
		g.compress(threshold)
		g.band = 0
	}
}

func (g *GKSummary) compress(threshold int) {
	if len(g.tuples) <= 2 || threshold <= 0 {
		return
	}
	for i := len(g.tuples) - 2; i >= 1; i-- {
		if g.tuples[i].g+g.tuples[i+1].g+g.tuples[i+1].delta <= threshold {
			g.tuples[i+1].g += g.tuples[i].g
			g.tuples = append(g.tuples[:i], g.tuples[i+1:]...)
		}
	}
}

func (g *GKSummary) Quantile(phi float64) (float64, error) {
	if g.n == 0 {
		return 0, ErrQuantileEmpty
	}
	if phi < 0 || phi > 1 {
		return 0, ErrQuantileBounds
	}
	threshold := int(math.Floor(2.0 * g.eps * float64(g.n)))
	desired := int(math.Ceil(phi * float64(g.n)))
	var rmin int
	for i := 0; i < len(g.tuples); i++ {
		rmin += g.tuples[i].g
		rmax := rmin + g.tuples[i].delta
		if rmax > desired+threshold/2 {
			if i == 0 {
				return g.tuples[0].v, nil
			}
			return g.tuples[i-1].v, nil
		}
	}
	return g.tuples[len(g.tuples)-1].v, nil
}
