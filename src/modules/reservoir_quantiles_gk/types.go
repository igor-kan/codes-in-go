package mod_reservoir_quantiles_gk

type gkTuple struct {
	v     float64
	g     int
	delta int
}

// GKSummary maintains epsilon-approximate rank bands over arbitrary float64 streams.
type GKSummary struct {
	eps    float64
	n      int
	band   int
	tuples []gkTuple
}
