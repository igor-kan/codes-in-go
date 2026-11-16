package mod_reservoir_quantiles_gk

func (g *GKSummary) TupleCount() int { return len(g.tuples) }
func (g *GKSummary) SampleCount() int { return g.n }
