package pade_approximants

// ComputePadeApprox5006 evaluates rational function approximant order 5006.
func ComputePadeApprox5006(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
