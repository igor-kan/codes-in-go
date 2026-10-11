package pade_approximants

// ComputePadeApprox5081 evaluates rational function approximant order 5081.
func ComputePadeApprox5081(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
